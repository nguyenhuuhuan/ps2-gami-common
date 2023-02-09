package notification

import (
	"bytes"
	"context"
	"encoding/json"
	"gitlab.id.vin/gami/gami-common/configs"
	"io/ioutil"
	"net/http"
)

// Const vin shop notification type
const (
	VinShopNtfBonusType   = "LOYALTY_BONUS"
	VinShopNtfRewardType  = "LOYALTY_REWARD"
	VinShopNtfNewTurnType = "LOYALTY_NEWTURN"
	VinShopNtfMaxTurnType = "LOYALTY_MAXTURN"
)

type Meta struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type MetaV2 struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Request struct {
	UserID     string   `json:"user_id"`
	Title      string   `json:"title"`
	Message    string   `json:"message"`
	Deeplink   string   `json:"deeplink"`
	Image      string   `json:"image"`
	Sound      bool     `json:"sound"`
	Vibration  bool     `json:"vibration"`
	ListUserID []string `json:"user_ids"`
}

type RequestNotiMTPV2 struct {
	AppID         string `json:"app_id"`
	TemplateID    string `json:"template_id"`
	CampaignID    string `json:"campaign_id"`
	SubjectID     string `json:"subject_id"`
	SendPush      bool   `json:"send_push"`
	ShowNotiStyle string `json:"show_noti_style"`
}

// SendNotificationRequest - request get token
type SendNotificationRequest struct {
	Token string
	Data  []Data `json:"data"`
}

// SendNotificationV2Request - request get token
type SendNotificationV2Request struct {
	AppID          string         `json:"app_id"`
	TemplateID     string         `json:"template_id"`
	Data           []DataNotiV2   `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type DataNotiV2 struct {
	UserID          string          `json:"user_id"`
	PersonalizeData PersonalizeData `json:"personalize_data"`
}

// Data notification
type Data struct {
	Type       string       `json:"type"`
	Recipients []Recipients `json:"recipients"`
	Data       ExtraData    `json:"data"`
}

// Recipients notification
type Recipients struct {
	Type     string `json:"type"`
	Identity string `json:"identity"`
}

// ExtraData notification
type ExtraData struct {
	Title     string `json:"title"`
	Message   string `json:"message"`
	DeepLink  string `json:"deep_link"`
	GroupID   int    `json:"group_id"`
	ChannelID string `json:"channel_id"`
}

type AdditionalData struct {
	Source string `json:"source"`
}

// PersonalizeData notification
type PersonalizeData struct {
	PushTitle   PushTitle   `json:"push_title"`
	PushMessage PushMessage `json:"push_message"`
	PushImage   PushImage   `json:"push_image"`
	Deeplink    Deeplink    `json:"deeplink"`
}

type PushTitle struct {
	Title string `json:"title"`
}

type PushMessage struct {
	Message string `json:"message"`
}

type Deeplink struct {
	Deeplink string `json:"deeplink"`
}

type PushImage struct {
	Image string `json:"image"`
}

// SendNotificationResponseV2 response get token
type SendNotificationV2Response struct {
	Data DataNotiV2 `json:"data"`
	Meta MetaV2     `json:"meta"`
}

// SendNotificationResponse response get token
type SendNotificationResponse struct {
	Meta Meta `json:"meta"`
}

func encodeSendNotificationRequest(_ context.Context, r *http.Request, req interface{}) error {
	request, _ := req.(*SendNotificationRequest)

	r.Header.Add("Authorization", request.Token)
	r.Header.Add("Content-Type", "application/json")

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)

	return nil
}

func decodeSendNotificationResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp SendNotificationResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func encodeSendNotificationV2Request(ctx context.Context, r *http.Request, req interface{}) error {
	request, _ := req.(*SendNotificationV2Request)

	r.Header.Add("Authorization", "Basic "+configs.AppConfig.Notification.ApiKey)
	r.Header.Add("Content-Type", "application/json")

	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(request)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buff)

	return nil
}

func decodeSendNotificationV2Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp SendNotificationV2Response
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type ExtraDataVinShop map[string]interface{}

type NtfVinShopRequest struct {
	UserID  string           `json:"user_id"`
	Title   string           `json:"title"`
	Message string           `json:"message"`
	Logo    string           `json:"logo"`
	ExData  ExtraDataVinShop `json:"ex_data"`
	Options *Options         `json:"options"`
	Channel int              `json:"channel"`
	Image   string           `json:"image"`
}

type SendNtfVinShopRequest struct {
	Topic    string           `json:"topic"`
	Title    string           `json:"title"`
	Message  string           `json:"message"`
	Logo     string           `json:"logo"`
	NotiType string           `json:"noti_type"`
	ExData   ExtraDataVinShop `json:"ex_data"`
	Options  *Options         `json:"options"`
	Channel  int              `json:"channel"` // 0, 1: ALL, 2: APP (iOS, Android), 3: SITE
	Image    string           `json:"image"`
}

// Options for sound options
type Options struct {
	IOS     *SoundInfo        `json:"ios"`
	Android *AndroidSoundInfo `json:"android"`
	Web     *SoundInfo        `json:"web"`
}

// SoundInfo contains sound
type SoundInfo struct {
	Sound string `json:"sound"`
}

// AndroidSoundInfo contains sound and channel ID
type AndroidSoundInfo struct {
	SoundInfo
	ChannelID string `json:"android_channel_id"`
}

type SendNtfVinShopResponse struct {
	Meta Meta `json:"meta"`
}

type CtaButton struct {
	Deeplink string `json:"deeplink"`
	Title    string `json:"title"`
}

type NtfVinShopV2Request struct {
	Message            string       `json:"message"`
	Title              string       `json:"title"`
	Topic              string       `json:"topic"`
	MessageType        string       `json:"message_type"`
	IsDataNotification bool         `json:"is_data_notification"`
	Deeplink           string       `json:"deeplink"`
	CtaButtons         []*CtaButton `json:"cta_buttons"`
	ImageURL           string       `json:"image_url"`
	ExtraData          interface{}  `json:"extra_data"`
}

type NtfVinShopV2Meta struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

type NtfVinShopV2Data struct {
	MessageID string `json:"message_id"`
}

type NtfVinShopV2Response struct {
	Meta NtfVinShopV2Meta `json:"meta"`
	Data NtfVinShopV2Data `json:"data"`
}

type MerchantInfo struct {
	MerchantID   int    `json:"merchant_id"`
	MerchantCode string `json:"merchant_code"`
}

type GetMerchantResponse struct {
	Meta Meta         `json:"meta"`
	Data MerchantInfo `json:"data"`
}
