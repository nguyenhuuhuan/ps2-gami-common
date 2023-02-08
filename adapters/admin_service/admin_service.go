package admin_service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"gitlab.id.vin/gami/ps2-gami-common/configs"
	"gitlab.id.vin/gami/ps2-gami-common/logger"

	"github.com/go-kit/kit/endpoint"
	transport "github.com/go-kit/kit/transport/http"
)

type AdminServiceAdapter interface {
	AdminServiceRequest(request *AdminServiceRequest) (*AdminServiceResponse, error)
}

type adminServiceAdapter struct {
	ctx          context.Context
	AdminService endpoint.Endpoint
}

func NewAdminServiceAdapter(ctx context.Context) AdminServiceAdapter {
	return &adminServiceAdapter{
		ctx:          ctx,
		AdminService: makeAdminServiceRequest(ctx),
	}
}

func makeAdminServiceRequest(ctx context.Context) endpoint.Endpoint {
	fullPath := configs.AppConfig.AdminService.BaseURL + configs.AppConfig.AdminService.Endpoint
	fullURL, err := url.Parse(fullPath)
	if err != nil {
		logger.Context(ctx).Errorf("[makeAdminServiceRequest] parsing url %v err: %v", fullPath, err)
	}

	return transport.NewClient(
		http.MethodPost, fullURL,
		encodeAdminServiceRequest,
		decodeAdminServiceResponse,
	).Endpoint()
}

func encodeAdminServiceRequest(ctx context.Context, r *http.Request, req interface{}) error {
	request, ok := req.(*AdminServiceRequest)
	if !ok {
		logger.Context(ctx).Errorf("[encodeAdminServiceRequest] type assertion failed.")
	}

	r.Header.Add("Authorization", "Basic "+configs.AppConfig.AdminService.BasicAuth)
	r.Header.Add("Content-Type", "application/json")

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeAdminServiceResponse(ctx context.Context, resp *http.Response) (interface{}, error) {
	var respData AdminServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		logger.Context(ctx).Errorf("[decodeAdminServiceResponse] decode failed, err: %v", err)
		return nil, err
	}

	if resp.StatusCode > 500 {
		logger.Context(ctx).Errorf("[decodeAdminServiceResponse] service unavailable")
		return &respData, errors.New("service unavailable")
	}

	return respData, nil
}

func (a *adminServiceAdapter) AdminServiceRequest(request *AdminServiceRequest) (*AdminServiceResponse, error) {
	r, err := a.AdminService(a.ctx, request)
	if err != nil {
		logger.Context(a.ctx).Errorf("[AdminServiceRequest] call admin service failed, req: %v, err: %v", request, err)
		return nil, err
	}

	resp, ok := r.(AdminServiceResponse)
	if r == nil || !ok {
		logger.Context(a.ctx).Errorf("[AdminServiceRequest] response is null")
		return nil, errors.New("response is null")
	}

	return &resp, nil
}
