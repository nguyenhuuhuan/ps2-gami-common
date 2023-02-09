package transaction_log

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"gitlab.id.vin/gami/gami-common/configs"
	"gitlab.id.vin/gami/gami-common/logger"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/endpoint"
)

// Const Path url
const (
	transactionPath = "/transaction-filters"
)

// Adapter handles all APIs for calling to users service.
type Adapter interface {
	GetTransactionLog(ctx context.Context, request *GetTransactionRequest) (*GetTransactionResponse, error)
}

type adapter struct {
	ctx                    context.Context
	GetTransactionEndpoint endpoint.Endpoint
}

// NewTransactionLogAdapter function
func NewTransactionLogAdapter(ctx context.Context) Adapter {
	return &adapter{
		ctx:                    ctx,
		GetTransactionEndpoint: makeGetTransaction(ctx),
	}
}

// Create to create campaign at mystery box service
func (a *adapter) GetTransactionLog(ctx context.Context, request *GetTransactionRequest) (*GetTransactionResponse, error) {
	r, err := a.GetTransactionEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	resp, ok := r.(GetTransactionResponse)
	if r == nil || !ok {
		return nil, errors.New("response is null")
	}
	return &resp, nil
}

func makeGetTransaction(ctx context.Context) endpoint.Endpoint {
	var (
		fullPath     = configs.AppConfig.Transaction.Endpoint + transactionPath
		fullURL, err = url.Parse(fullPath)
	)

	if err != nil {
		logger.Context(ctx).Errorf("Parsing url: %v", fullPath)
	}

	return httptransport.NewClient(
		http.MethodPost, fullURL,
		encodeGetTransactionLogRequest,
		decodeGetTransactionLogRequest,
		httptransport.ClientBefore(setRequestHeader()),
	).Endpoint()
}

func setRequestHeader() httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		{
			r.Header.Set("X-Api-Key", configs.AppConfig.Transaction.Authorization)
			r.Header.Set("Content-Type", "application/json")
		}
		return ctx
	}
}
