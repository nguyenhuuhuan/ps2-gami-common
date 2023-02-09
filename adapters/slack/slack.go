package slack

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"gitlab.id.vin/gami/gami-common/configs"
	"gitlab.id.vin/gami/gami-common/logger"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Const Path url
const (
	UpdateRulePath = "/rule-status"
)

// SlackAdapter handles all APIs for calling to users service.
type SlackAdapter interface {
	Send(request *Request) error
}

type slackAdapter struct {
	ctx           context.Context
	SlackEndpoint endpoint.Endpoint
}

// NewRuleAdapter returns a new instance of SlackAdapter.
func NewSlackAdapter(ctx context.Context) SlackAdapter {
	return &slackAdapter{
		ctx:           ctx,
		SlackEndpoint: makeSlackRequest(ctx),
	}
}

func makeSlackRequest(ctx context.Context) endpoint.Endpoint {

	var (
		fullPath     = configs.AppConfig.Slack.Endpoint + "/" + configs.AppConfig.Slack.Channel + "/" + configs.AppConfig.Slack.Token
		fullURL, err = url.Parse(fullPath)
	)

	if err != nil {
		logger.Context(ctx).Errorf("Parsing url: %v", fullPath)
	}

	return httptransport.NewClient(
		http.MethodPost, fullURL,
		encodeSlackRequest,
		decodeSlackRequest,
	).Endpoint()
}

func (a *slackAdapter) Send(request *Request) error {
	r, err := a.SlackEndpoint(a.ctx, request)
	if err != nil {
		return err
	}

	if r == nil {
		return errors.New("response is null")
	}

	return nil
}
