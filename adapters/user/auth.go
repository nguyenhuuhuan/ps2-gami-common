package user

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"gitlab.id.vin/gami/gami-common/configs"
	"gitlab.id.vin/gami/gami-common/dtos"
	"gitlab.id.vin/gami/gami-common/logger"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/patrickmn/go-cache"
)

// Const Path url
const (
	TokenPath       = "/token"
	InternalProfile = "/internal/profile"
	ListProfile     = "/internal/search"
)

// AuthAdapter handles all APIs for calling to users service.
type AuthAdapter interface {
	GetToken(request *TokenRequest) (*GetTokenResponse, error)
	GetUserProfile(request *ProfileRequest) (*ProfileResponse, error)
	GetListUserProfile(request *ListUserRequest) (*ListUserResponse, error)
	SetCacheToken(key, token string) bool
	GetCacheToken(key string) (string, bool)
}

type authAdapter struct {
	ctx                    context.Context
	GetProfileEndpoint     endpoint.Endpoint
	GetTokenEndpoint       endpoint.Endpoint
	GetListProfileEndpoint endpoint.Endpoint
	localCacheAdapter      *cache.Cache
}

// NewAuthAdapter returns a new instance of AuthAdapter.
func NewAuthAdapter(ctx context.Context) AuthAdapter {
	return &authAdapter{
		ctx:                    ctx,
		GetTokenEndpoint:       makeToken(ctx),
		GetProfileEndpoint:     makeGetUserProfile(ctx),
		GetListProfileEndpoint: makeGetListProfile(ctx),
		localCacheAdapter:      cache.New(50*time.Minute, 1*time.Minute),
	}
}

func makeGetUserProfile(ctx context.Context) endpoint.Endpoint {
	var (
		fullPath     = configs.AppConfig.User.AccountEndpoint + InternalProfile
		fullURL, err = url.Parse(fullPath)
	)

	if err != nil {
		logger.Context(ctx).Errorf("cannot parse url %v", err)
	}

	return httptransport.NewClient(
		http.MethodGet, fullURL,
		encodeGetUserProfileRequest,
		decodeGetUserProfileResponse,
		httptransport.SetClient(&http.Client{
			Timeout: 10 * time.Second,
		}),
	).Endpoint()
}

func makeGetListProfile(ctx context.Context) endpoint.Endpoint {
	var (
		fullPath     = configs.AppConfig.User.ProfileEndpoint + ListProfile
		fullURL, err = url.Parse(fullPath)
	)

	if err != nil {
		logger.Context(ctx).Errorf("cannot parse url %v", err)
	}

	return httptransport.NewClient(
		http.MethodGet, fullURL,
		encodeListProfileRequest,
		decodeListProfileResponse,
		httptransport.SetClient(&http.Client{
			Timeout: 30 * time.Second,
		}),
	).Endpoint()
}

func (a *authAdapter) GetUserProfile(request *ProfileRequest) (*ProfileResponse, error) {
	if request.Token == "" {
		data, exists := a.localCacheAdapter.Get(configs.AppConfig.User.UserScope)
		if !exists {
			tokenReq := TokenRequest{
				Scope:     configs.AppConfig.User.UserScope,
				GrantType: "client_credentials",
			}
			tokenResp, err := a.GetToken(&tokenReq)
			if err != nil || tokenResp == nil || tokenResp.AccessToken == "" {
				return nil, errors.New(fmt.Sprintf("Get token error %v", err))
			}

			token := "Bearer " + tokenResp.AccessToken
			a.SetCacheToken(configs.AppConfig.User.UserScope, token)
			request.Token = token
		} else {
			token, _ := data.(string)
			request.Token = token
		}
	}

	r, err := a.GetProfileEndpoint(a.ctx, request)
	if err != nil {
		logger.Context(context.Background()).Errorf("get profile error %v", err)
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, dtos.NewAppError(dtos.RequestTimeoutError)
		}
		return nil, err
	}

	resp, ok := r.(ProfileResponse)
	if r == nil || !ok || resp.Meta.Code != http.StatusOK {
		return nil, dtos.NewAppError(dtos.UserNotFoundError)
	}
	return &resp, nil
}

func (a *authAdapter) GetListUserProfile(request *ListUserRequest) (*ListUserResponse, error) {
	if request.Token == "" {
		data, exists := a.localCacheAdapter.Get(configs.AppConfig.User.UserScope)
		if !exists {
			tokenReq := TokenRequest{
				Scope:     "internal.profile",
				GrantType: "client_credentials",
			}

			tokenResp, err := a.GetToken(&tokenReq)
			if err != nil || tokenResp == nil || tokenResp.AccessToken == "" {
				return nil, errors.New(fmt.Sprintf("Get token error %v", err))
			}
			_ = a.SetCacheToken(configs.AppConfig.User.UserScope, "Bearer "+tokenResp.AccessToken)
			request.Token = "Bearer " + tokenResp.AccessToken
		} else {
			token, _ := data.(string)
			request.Token = token
		}
	}

	r, err := a.GetListProfileEndpoint(a.ctx, request)
	if err != nil {
		logger.Context(context.Background()).Errorf("get profile error %v", err)
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, dtos.NewAppError(dtos.RequestTimeoutError)
		}
		return nil, err
	}

	resp, ok := r.(ListUserResponse)
	if r == nil || !ok || resp.Meta.Code != http.StatusOK {
		return nil, dtos.NewAppError(dtos.UserNotFoundError)
	}
	return &resp, nil
}

func makeToken(ctx context.Context) endpoint.Endpoint {
	var (
		fullPath     = configs.AppConfig.User.OauthEndpoint + TokenPath
		fullURL, err = url.Parse(fullPath)
	)

	if err != nil {
		logger.Context(ctx).Errorf("cannot parse url %v", err)
	}

	return httptransport.NewClient(
		http.MethodPost, fullURL,
		encodeTokenRequest,
		decodeTokenResponse,
		httptransport.ClientBefore(setRequestHeader()),
		httptransport.SetClient(&http.Client{
			Timeout: 10 * time.Second,
		}),
	).Endpoint()
}

func (a *authAdapter) GetToken(request *TokenRequest) (*GetTokenResponse, error) {
	r, err := a.GetTokenEndpoint(a.ctx, request)
	if err != nil {
		logger.Context(a.ctx).Errorf("call get token error %v", err)
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil, dtos.NewAppError(dtos.RequestTimeoutError)
		}
		return nil, err
	}

	resp, ok := r.(GetTokenResponse)
	if r == nil || !ok || resp.AccessToken == "" {
		return nil, dtos.NewAppError(dtos.GetTokenError)
	}
	return &resp, nil
}

func setRequestHeader() httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		{
			r.Header.Set("Authorization", "Basic "+configs.AppConfig.User.IntegrationHeader.Authorization)
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		return ctx
	}
}

func (a *authAdapter) SetCacheToken(key, token string) bool {
	if token == "" {
		return false
	}

	a.localCacheAdapter.Set(key, token, 50*time.Minute)
	return true
}

func (a *authAdapter) GetCacheToken(key string) (string, bool) {
	data, exists := a.localCacheAdapter.Get(key)
	if !exists {
		return "", exists
	}

	token, ok := data.(string)
	return token, ok
}
