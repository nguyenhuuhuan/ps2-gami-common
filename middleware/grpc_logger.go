package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"gitlab.id.vin/gami/ps2-gami-common/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// LoggerUnaryInterceptor returns a new unary server interceptors that add to the context.
func LoggerUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		start := time.Now()
		newContext := logger.SetRqIDToCtx(ctx)

		resp, respErr := handler(newContext, req)
		md, _ := metadata.FromIncomingContext(newContext)
		end := time.Now()
		latency := end.Sub(start)

		fields := logger.Fields{
			"rqID":       logger.GetRqIDFromCtx(newContext),
			"method":     info.FullMethod,
			"time":       end.Unix(),
			"authority":  md[":authority"],
			"user-agent": md["user-agent"],
			"exec_time":  float64(latency) / 1000000,
		}

		data, err := json.Marshal(req)
		if err == nil {
			buffer := new(bytes.Buffer)
			err = json.Compact(buffer, data)
			fields["data"] = buffer
		}

		if respErr != nil {
			logger.WithFields(fields).Errorf("Error %v", respErr.Error())
		} else {
			logger.WithFields(fields).Infof("Success")
		}

		return resp, err
	}
}

var (
	defaultOptions = &options{
		recoveryHandlerFunc: nil,
	}
)

type options struct {
	recoveryHandlerFunc RecoveryHandlerFuncContext
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

type Option func(*options)

// WithRecoveryHandler customizes the function for recovering from a panic.
func WithRecoveryHandler(f RecoveryHandlerFunc) Option {
	return func(o *options) {
		o.recoveryHandlerFunc = RecoveryHandlerFuncContext(func(ctx context.Context, p interface{}) error {
			return f(p)
		})
	}
}

// WithRecoveryHandlerContext customizes the function for recovering from a panic.
func WithRecoveryHandlerContext(f RecoveryHandlerFuncContext) Option {
	return func(o *options) {
		o.recoveryHandlerFunc = f
	}
}

// RecoveryHandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type RecoveryHandlerFunc func(p interface{}) (err error)

// RecoveryHandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerFuncContext func(ctx context.Context, p interface{}) (err error)

// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = recoverFrom(ctx, r, o.recoveryHandlerFunc)
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false
		return resp, err
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for panic recovery.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = recoverFrom(stream.Context(), r, o.recoveryHandlerFunc)
			}
		}()

		err = handler(srv, stream)
		panicked = false
		return err
	}
}

func recoverFrom(ctx context.Context, p interface{}, r RecoveryHandlerFuncContext) error {
	if r == nil {
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, p)
}
