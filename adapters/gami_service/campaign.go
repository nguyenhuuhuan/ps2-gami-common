package gami_service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	"gitlab.id.vin/gami/ps2-gami-common/dtos"
	gami_protobuf "gitlab.id.vin/gami/ps2-gami-proto/pb"
)

type GetCampaignRequest struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

type GetCampaignUserRequest struct {
	ID     int64  `json:"id"`
	Code   string `json:"code"`
	UserID int64  `json:"user_id"`
}

type GetCampaignResponse struct {
	Meta dtos.Meta `json:"meta"`
	Data *Campaign `json:"data"`
}

type ListCampaignsRequest struct {
	Type      string `json:"type"`
	UiType    string `json:"ui_type"`
	DateStart int64  `json:"date_start"`
	DateEnd   int64  `json:"date_end"`
	Offset    int64  `json:"offset"`
	Limit     int64  `json:"limit"`
	IsActive  bool   `json:"is_active"`
}

type ListCampaignsResponse struct {
	Meta dtos.Meta   `json:"meta"`
	Data []*Campaign `json:"data"`
}

type GetListCampaignByIDsRequest struct {
	IDs []int64 `json:"ids"`
}

type Campaign struct {
	ID          int64   `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	Banner      string  `json:"banner"`
	DateStart   int64   `json:"date_start"`
	DateEnd     int64   `json:"date_end"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	Config      []byte  `json:"config"`
	CreatedAt   int64   `json:"created_at"`
	Rules       []*Rule `json:"rules"`
	AppType     string  `json:"app_type"`
}

// Rule struct
type Rule struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	RuleID    int64  `json:"rule_id"`
	Quantity  int64  `json:"quantity"`
	Icon      string `json:"icon"`
	Deeplink  string `json:"deeplink"`
	Status    string `json:"status"`
	Content   string `json:"content"`
	IsShown   bool   `json:"is_shown"`
	CreatedAt int64  `json:"created_at"`
}

// AcceptChallengeRequest
type AcceptChallengeRequest struct {
	UserID     string `json:"user_id"`
	CampaignID int64  `json:"campaign_id"`
}

// AcceptChallengeResponse grpc response.
type AcceptChallengeResponse struct {
	Meta dtos.Meta `json:"meta"`
}

type GetCampaignByRuleID struct {
	RuleID int64 `json:"rule_id"`
}

type RewardAmountStatisticRequest struct {
	UserID     string `json:"user_id"`
	CampaignID int64  `json:"campaign_id"`
	RewardID   int64  `json:"reward_id"`
	FromDate   int64  `json:"from_date"`
	ToDate     int64  `json:"to_date"`
}

// AcceptChallengeResponse grpc response.
type RewardAmountStatisticResponse struct {
	Meta   dtos.Meta `json:"meta"`
	Amount int64     `json:"amount"`
}

func decodeGetCampaignResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetCampaignResponse)

	if r.Meta.Status != http.StatusOK {
		return &GetCampaignResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	campaignData := Campaign{}
	_ = copier.Copy(&campaignData, r.Data)

	data := GetCampaignResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: &campaignData,
	}

	return &data, nil

}

func encodeGetCampaignRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, _ := req.(GetCampaignRequest)
	return &gami_protobuf.GetCampaignRequest{
		ID:   r.ID,
		Code: r.Code,
	}, nil
}

func encodeGetCampaignUserRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, _ := req.(GetCampaignUserRequest)
	return &gami_protobuf.GetCampaignUserRequest{
		ID:     r.ID,
		Code:   r.Code,
		UserID: r.UserID,
	}, nil
}

func encodeListCampaignsRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, _ := req.(*ListCampaignsRequest)
	return &gami_protobuf.GetListCampaignRequest{
		Type:      r.Type,
		UiType:    r.UiType,
		DateStart: r.DateStart,
		DateEnd:   r.DateEnd,
		Offset:    r.Offset,
		Limit:     r.Limit,
		IsActive:  r.IsActive,
	}, nil
}

func encodeAcceptChallengeRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*AcceptChallengeRequest)
	if !ok {
		return nil, errors.New("cast error")
	}
	return &gami_protobuf.AcceptChallengeRequest{
		UserID:     r.UserID,
		CampaignID: r.CampaignID,
	}, nil
}

func encodeGetCampaignByRuleIDRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*GetCampaignByRuleID)
	if !ok {
		return nil, errors.New("cast error")
	}
	return &gami_protobuf.GetCampaignByRuleIDRequest{
		RuleID: r.RuleID,
	}, nil
}

func decodeAcceptChallengeResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, ok := resp.(*gami_protobuf.AcceptChallengeResponse)
	if !ok {
		return nil, errors.New("cast error")
	}
	if r.Meta.Status != http.StatusOK {
		return &AcceptChallengeResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}
	return &AcceptChallengeResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
	}, nil
}

func decodeListCampaignsResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetListCampaignResponse)
	if r.Meta.Status != http.StatusOK {
		return &ListCampaignsResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	var data = make([]*Campaign, len(r.Data))
	for index, item := range r.Data {
		var campaign Campaign
		_ = copier.Copy(&campaign, item)
		data[index] = &campaign
	}

	return &ListCampaignsResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: data,
	}, nil
}

func encodeGetListCampaignByIDsRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, _ := req.(*GetListCampaignByIDsRequest)
	return &gami_protobuf.GetListCampaignByIDsRequest{
		IDs: r.IDs,
	}, nil
}

func encodeRewardAmountStatisticRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*RewardAmountStatisticRequest)
	if !ok {
		return nil, errors.New("cast error")
	}
	return &gami_protobuf.RewardAmountStatisticRequest{
		UserID:     r.UserID,
		CampaignID: r.CampaignID,
		RewardID:   r.RewardID,
		FromDate:   r.FromDate,
		ToDate:     r.ToDate,
	}, nil
}

func decodeRewardAmountStatisticResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, ok := resp.(*gami_protobuf.RewardAmountStatisticResponse)
	if !ok {
		return nil, errors.New("cast error")
	}
	if r.Meta.Status != http.StatusOK {
		return &RewardAmountStatisticResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}
	return &RewardAmountStatisticResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Amount: r.Amount,
	}, nil
}

func decodeGetCampaignByRuleIDResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, ok := resp.(*gami_protobuf.GetCampaignResponse)
	if !ok {
		return nil, errors.New("cast error")
	}
	if r.Meta.Status != http.StatusOK {
		return &GetCampaignResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}
	campaignData := Campaign{}
	_ = copier.Copy(&campaignData, r.Data)

	data := GetCampaignResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: &campaignData,
	}

	return &data, nil
}
