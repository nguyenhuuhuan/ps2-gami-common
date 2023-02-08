package makerchecker

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"gitlab.id.vin/gami/go-agent/v3/newrelic"

	"gitlab.id.vin/gami/ps2-gami-common/configs"
	"gitlab.id.vin/gami/ps2-gami-common/logger"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Const Path url
const (
	MakerPath = "/changes"
)

// Adapter handles all APIs for calling to maker checker service.
type Adapter interface {
	CreateRequest(request *MakerRequest) (*MakerResponse, error)
}

type makerCheckerAdapter struct {
	ctx             context.Context
	MakerNewRequest endpoint.Endpoint
}

// NewMakerCheckerAdapter returns a new instance of MakerCheckerAdapter.
func NewMakerCheckerAdapter(ctx context.Context) Adapter {
	return &makerCheckerAdapter{
		ctx:             ctx,
		MakerNewRequest: makeMakerRequest(ctx),
	}
}

func makeMakerRequest(ctx context.Context) endpoint.Endpoint {
	var tracingClient = http.DefaultClient
	tracingClient.Transport = newrelic.NewRoundTripper(tracingClient.Transport)

	var (
		fullPath     = configs.AppConfig.MakerChecker.Endpoint + MakerPath
		fullURL, err = url.Parse(fullPath)
	)

	if err != nil {
		logger.Context(ctx).Errorf("Parsing url: %v", fullPath)
	}

	return httptransport.NewClient(
		http.MethodPost, fullURL,
		encodeMakerRequest,
		decodeMakerResponse,
		httptransport.ClientBefore(setRequestHeader()),
		httptransport.SetClient(tracingClient),
	).Endpoint()
}

func (a *makerCheckerAdapter) CreateRequest(request *MakerRequest) (*MakerResponse, error) {
	r, err := a.MakerNewRequest(a.ctx, request)
	if err != nil {
		return nil, err
	}

	resp, ok := r.(MakerResponse)
	if r == nil || !ok {
		return nil, errors.New("response is null")
	}
	return &resp, nil
}

func setRequestHeader() httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		{
			r.Header.Set("Content-Type", "application/json")
			r.Header.Set("Authorization", "Basic "+configs.AppConfig.MakerChecker.Authorization)
		}
		return ctx
	}
}
