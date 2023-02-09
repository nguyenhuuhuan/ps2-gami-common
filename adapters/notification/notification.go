package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.id.vin/gami/gami-common/adapters/cache"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"gitlab.id.vin/gami/gami-common/adapters/user"
	"gitlab.id.vin/gami/gami-common/configs"
	"gitlab.id.vin/gami/gami-common/logger"
)

// Const Path url
const (
	SendPath                = "/send"
	VinShopAllowChannel int = 2
	VinShopAnnouncement     = "ANNOUNCEMENT"
	VinShopPromotion        = "PROMOTION"
	VinShopNews             = "NEWS"
)

// Adapter handles all APIs for calling to users service.
type Adapter interface {
	SendNtfVinShop(ctx context.Context, request *NtfVinShopRequest) (*SendNtfVinShopResponse, error)
	SendNtfVinShopV2(ctx context.Context, request *NtfVinShopRequest) (*NtfVinShopV2Response, error)
	SendNotification(ctx context.Context, request *Request) (*SendNotificationResponse, error)
	SendNotificationV2(ctx context.Context, request *Request) (*SendNotificationV2Response, error)
}

type adapter struct {
	ctx                        context.Context
	authAdapter                user.AuthAdapter
	localCacheAdapter          cache.LocalCacheAdapter
	merchantInfoEndpoint       endpoint.Endpoint
	sendNtfVinShopEndpoint     endpoint.Endpoint
	sendNtfVinShopV2Endpoint   endpoint.Endpoint
	sendNotificationEndpoint   endpoint.Endpoint
	sendNotificationV2Endpoint endpoint.Endpoint
}

// NewAdapter returns a new instance of Adapter.
func NewAdapter(ctx context.Context, authAdapter user.AuthAdapter) Adapter {
	localCacheAdapter := cache.NewLocalCacheAdapter()
	return &adapter{
		ctx:                        ctx,
		authAdapter:                authAdapter,
		localCacheAdapter:          localCacheAdapter,
		sendNtfVinShopEndpoint:     makeSendNtfVinShop(ctx),
		sendNotificationEndpoint:   makeSendNoti(ctx),
		sendNotificationV2Endpoint: makeSendNotiV2(ctx),
		sendNtfVinShopV2Endpoint:   makeSendNtfVinShopV2(ctx),
		merchantInfoEndpoint:       getMerchantInfoEndpoint(ctx),
	}
}

func makeSendNoti(ctx context.Context) endpoint.Endpoint {
	var (
		fullPath     = configs.AppConfig.Notification.Endpoint + SendPath
		fullURL, err = url.Parse(fullPath)
	)

	if err != nil {
		logger.Context(ctx).Errorf("cannot parse url %v", err)
	}

	return httptransport.NewClient(
		http.MethodPost, fullURL,
		encodeSendNotificationRequest,
		decodeSendNotificationResponse,
	).Endpoint()
}

func makeSendNotiV2(ctx context.Context) endpoint.Endpoint {
	var (
		fullPath     = configs.AppConfig.Notification.EndpointV2 + SendPath
		fullUrl, err = url.Parse(fullPath)
	)
	if err != nil {
		logger.Context(ctx).Errorf("cannot parse url %v", err)
	}

	return httptransport.NewClient(
		http.MethodPost, fullUrl,
		encodeSendNotificationV2Request,
		decodeSendNotificationV2Response,
	).Endpoint()
}

func (a *adapter) SendNotificationV2(ctx context.Context, request *Request) (*SendNotificationV2Response, error) {

	if len(request.ListUserID) == 0 {
		request.ListUserID = append(request.ListUserID, request.UserID)
	}

	var dataNotiV2 = []DataNotiV2{}
	for _, ID := range request.ListUserID {
		dataNotiV2 = append(dataNotiV2, DataNotiV2{
			UserID: ID,
			PersonalizeData: PersonalizeData{
				PushTitle: PushTitle{
					Title: request.Title,
				},
				PushMessage: PushMessage{
					Message: request.Message,
				},
				PushImage: PushImage{
					Image: request.Image,
				},
				Deeplink: Deeplink{
					Deeplink: request.Deeplink,
				},
			},
		})
	}
	var additionalData = AdditionalData{
		Source: "poke",
	}

	req := SendNotificationV2Request{
		AppID:          configs.AppConfig.Notification.AppID,
		TemplateID:     configs.AppConfig.Notification.TemplateID,
		Data:           dataNotiV2,
		AdditionalData: additionalData,
	}
	json, _ := json.Marshal(req)
	logger.Context(ctx).Infof("[SendNotificationV2] request %v:", string(json))
	resp, err := a.sendNotificationV2(&req)
	return resp, err
}

func (a *adapter) sendNotificationV2(request *SendNotificationV2Request) (*SendNotificationV2Response, error) {
	resp, err := a.sendNotificationV2Endpoint(a.ctx, request)
	if err != nil {
		return nil, err
	}

	data, ok := resp.(SendNotificationV2Response)
	if resp == nil || !ok {
		return nil, errors.New("response is null")
	}

	return &data, nil

}

func (a *adapter) SendNotification(ctx context.Context, request *Request) (*SendNotificationResponse, error) {
	token, err := a.getToken(configs.AppConfig.User.NotificationScope, a.authAdapter)
	if err != nil {
		return nil, err
	}

	if len(request.ListUserID) == 0 {
		request.ListUserID = append(request.ListUserID, request.UserID)
	}

	var recipients = []Recipients{}

	for _, ID := range request.ListUserID {
		recipients = append(recipients, Recipients{
			Type:     "user_id",
			Identity: ID,
		})
	}

	notificationData := Data{
		Type:       "notification",
		Recipients: recipients,
		Data: ExtraData{
			Title:     request.Title,
			Message:   request.Message,
			DeepLink:  request.Deeplink,
			ChannelID: configs.AppConfig.Notification.ChannelID,
			GroupID:   configs.AppConfig.Notification.GroupID,
		},
	}

	req := SendNotificationRequest{
		Data:  []Data{notificationData},
		Token: token,
	}

	resp, err := a.sendNotification(&req)
	return resp, err
}

func (a *adapter) sendNotification(request *SendNotificationRequest) (*SendNotificationResponse, error) {
	r, err := a.sendNotificationEndpoint(a.ctx, request)
	if err != nil {
		return nil, err
	}

	resp, ok := r.(SendNotificationResponse)
	if r == nil || !ok {
		return nil, errors.New("response is null")
	}

	return &resp, nil
}

func (a *adapter) getUserProfile(userID string) (*user.ProfileResponse, error) {
	var profile *user.ProfileResponse

	token, err := a.getToken(configs.AppConfig.User.UserScope, a.authAdapter)
	if err != nil {
		return nil, err
	}

	request := user.ProfileRequest{
		UserID: userID,
		Type:   "USER_ID",
		Token:  token,
	}

	profile, err = a.authAdapter.GetUserProfile(&request)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (a *adapter) getToken(scope string, adapter user.AuthAdapter) (string, error) {
	request := user.TokenRequest{
		Scope:     scope,
		GrantType: "client_credentials",
	}

	token, ok := a.authAdapter.GetCacheToken(scope)
	if !ok || token == "" {
		response, err := adapter.GetToken(&request)
		if err != nil || response == nil {
			return "", err
		}

		ok = a.authAdapter.SetCacheToken(scope, "Bearer "+response.AccessToken)
		if !ok {
			logger.Errorf("Set cache error")
		}
		token = "Bearer " + response.AccessToken
		return token, nil
	}

	return token, nil
}

// Send Ntf to vin-shop app
func (a *adapter) SendNtfVinShop(ctx context.Context, request *NtfVinShopRequest) (*SendNtfVinShopResponse, error) {
	if request.Channel != VinShopAllowChannel {
		msg := fmt.Sprintf("not allow channel with value %v", request.Channel)
		logger.Context(ctx).Errorf("[SendNtfVinShop] %v", msg)

		return nil, errors.New(msg)
	}

	ntfReq := SendNtfVinShopRequest{
		Topic:    configs.AppConfig.VinShop.NtfTopicPrefix + request.UserID,
		Title:    request.Title,
		Message:  request.Message,
		Logo:     request.Logo,
		NotiType: configs.AppConfig.VinShop.NtfType,
		ExData:   request.ExData,
		Options:  request.Options,
		Channel:  request.Channel,
		Image:    request.Image,
	}

	r, err := a.sendNtfVinShopEndpoint(a.ctx, &ntfReq)
	if err != nil {
		logger.Context(ctx).Errorf("[SendNtfVinShop] send request failed: %v", err)
		return nil, err
	}

	resp, ok := r.(SendNtfVinShopResponse)
	if r == nil || !ok {
		logger.Context(ctx).Errorf("[SendNtfVinShop] response is null")
		return nil, errors.New("response is null")
	}

	return &resp, nil
}

func makeSendNtfVinShop(ctx context.Context) endpoint.Endpoint {
	fullPath := configs.AppConfig.VinShop.NtfBaseURL + configs.AppConfig.VinShop.NtfSendEndpoint

	fullURL, err := url.Parse(fullPath)
	if err != nil {
		logger.Context(ctx).Errorf("[makeSendNtfVinShop] cannot parse url %v", err)
	}

	return httptransport.NewClient(
		http.MethodPost, fullURL,
		encodeSendNtfVinShopRequest,
		decodeSendNtfVinShopResponse,
	).Endpoint()
}

func encodeSendNtfVinShopRequest(ctx context.Context, r *http.Request, req interface{}) error {
	request, _ := req.(*SendNtfVinShopRequest)

	// generate basic auth
	basicAuth := fmt.Sprintf("Basic %v", configs.AppConfig.VinShop.NtfAuth)

	// Set header
	r.Header.Add("Authorization", basicAuth)
	r.Header.Add("Content-Type", "application/json")

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(&buf)

	return nil
}

func decodeSendNtfVinShopResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp SendNtfVinShopResponse

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *adapter) SendNtfVinShopV2(ctx context.Context, request *NtfVinShopRequest) (*NtfVinShopV2Response, error) {
	var data GetMerchantResponse
	var key = fmt.Sprintf("merchant:userID:%v", request.UserID)

	err := a.localCacheAdapter.Get(ctx, key, &data)
	if err != nil {
		resp, err := a.merchantInfoEndpoint(ctx, request.UserID)
		if err != nil {
			logger.Context(ctx).Errorf("get merchant information for userid %v error %v", request.UserID, err)
			return nil, err
		}
		r, ok := resp.(*GetMerchantResponse)
		if r == nil || !ok {
			return nil, errors.New("cast type fail")
		}
		if r.Meta.Code == http.StatusOK {
			_ = a.localCacheAdapter.Set(ctx, key, r, 24*time.Hour)
		} else {
			logger.Context(ctx).Errorf("get merchant information for user_id %v has invalid resp : %v", request.UserID, r)
			return nil, errors.New("invalid resp")
		}
	}

	ntfReq := NtfVinShopV2Request{
		Topic:       configs.AppConfig.VinShop.NtfTopicPrefix + data.Data.MerchantCode,
		Title:       request.Title,
		Message:     request.Message,
		MessageType: configs.AppConfig.VinShop.NtfType,
	}

	r, err := a.sendNtfVinShopV2Endpoint(a.ctx, &ntfReq)
	if err != nil {
		logger.Context(ctx).Errorf("[SendNtfVinShopv2] send request failed: %v", err)
		return nil, err
	}

	resp, ok := r.(NtfVinShopV2Response)
	if r == nil || !ok {
		logger.Context(ctx).Errorf("[SendNtfVinShopv2] response is null")
		return nil, errors.New("response is null")
	}

	return &resp, nil

}

func makeSendNtfVinShopV2(ctx context.Context) endpoint.Endpoint {
	fullPath := configs.AppConfig.VinShop.NtfBaseURL + configs.AppConfig.VinShop.NtfSendEndpoint

	fullURL, err := url.Parse(fullPath)
	if err != nil {
		logger.Context(ctx).Errorf("[makeSendNtfVinShopV2] cannot parse url %v", err)
	}

	return httptransport.NewClient(
		http.MethodPost, fullURL,
		encodeSendNtfVinShopV2Request,
		decodeSendNtfVinShopV2Response,
	).Endpoint()
}

func encodeSendNtfVinShopV2Request(ctx context.Context, r *http.Request, req interface{}) error {
	request, _ := req.(*NtfVinShopV2Request)

	basicAuth := fmt.Sprintf("Basic %v", configs.AppConfig.VinShop.NtfAuth)
	r.Header.Add("Authorization", basicAuth)
	r.Header.Add("Content-Type", "application/json")

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(&buf)

	return nil
}

func decodeSendNtfVinShopV2Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp NtfVinShopV2Response

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func getMerchantInfoEndpoint(ctx context.Context) endpoint.Endpoint {
	fullPath := configs.AppConfig.MerchantInformation.BaseURL + configs.AppConfig.MerchantInformation.Endpoint
	fullURL, err := url.Parse(fullPath)
	if err != nil {
		logger.Context(ctx).Errorf("cannot parse url %v error %v", fullPath, err)
	}

	return httptransport.NewClient(
		http.MethodGet, fullURL, encodeMerchantInformation, decodeMerchantInformation,
		httptransport.SetClient(&http.Client{
			Timeout: 30 * time.Second,
		}),
		httptransport.ClientBefore(httptransport.SetRequestHeader("Authorization", configs.AppConfig.MerchantInformation.BasicAuth),
			httptransport.SetRequestHeader("Content-Type", "application/json")),
	).Endpoint()
}

func encodeMerchantInformation(ctx context.Context, request *http.Request, req interface{}) error {
	v, ok := req.(string)
	if !ok {
		logger.Context(ctx).Errorf("request %v is not valid type", req)
		return errors.New("invalid request")
	}
	q := request.URL.Query()
	q.Add("vinid_user_id", v)
	request.URL.RawQuery = q.Encode()
	return nil
}

func decodeMerchantInformation(ctx context.Context, resp *http.Response) (interface{}, error) {
	var result GetMerchantResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Context(ctx).Errorf("decode resp %v error : %v", resp.Body, err)
		return nil, err
	}
	return &result, nil
}
