package gami_service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	gami_protobuf "gitlab.id.vin/gami/gami-proto/pb"
	"gitlab.id.vin/gami/ps2-gami-common/dtos"
)

type GetRewardResponse struct {
	Meta dtos.Meta `json:"meta"`
	Data *Reward   `json:"data"`
}

type ListRewardResponse struct {
	Meta dtos.Meta `json:"meta"`
	Data []*Reward `json:"data"`
}

type Reward struct {
	ID          int64    `json:"id"`
	Code        string   `json:"code"`
	Status      string   `json:"status"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Deeplink    string   `json:"deeplink"`
	Config      []byte   `json:"config"`
	Voucher     *Voucher `json:"voucher"`
}

type Voucher struct {
	Code  string `json:"code"`
	Image string `json:"image"`
}

type GetRewardPoolRequest struct {
	CampaignID   int64  `json:"campaign_id"`
	CampaignCode string `json:"campaign_code"`
}

type GetListRewardByIDsRequest struct {
	IDs []int64 `json:"ids"`
}

type RewardPool struct {
	ID         int64   `json:"id"`
	RewardId   int64   `json:"reward_id"`
	Reward     *Reward `json:"reward"`
	AmountLeft int64   `json:"amount_left"`
	UnitValue  int64   `json:"unit_value"`
	Status     string  `json:"status"`
}

type GetRewardPoolResponse struct {
	Meta dtos.Meta     `json:"meta"`
	Data []*RewardPool `json:"data"`
}

type GetRewardsByCampaignIDResponse struct {
	Meta dtos.CursorMeta `json:"meta"`
	Data []*RewardPool   `json:"data"`
}

// RedeemRequest struct
type RedeemRequest struct {
	CampaignID   int64           `json:"campaign_id"`
	UserID       string          `json:"user_id"`
	RewardPoolID int64           `json:"reward_pool_id"`
	Quantity     int             `json:"quantity"`
	PoolID       int64           `json:"pool_id"`
	ExtraData    json.RawMessage `json:"extra_data,omitempty"`
}

// RedeemResponse struct define
type RedeemResponse struct {
	Meta dtos.Meta   `json:"meta"`
	Data *RedeemData `json:"data"`
}

// RedeemData struct define
type RedeemData struct {
	RewardPoolID int64   `json:"reward_pool_id"`
	Reward       *Reward `json:"reward"`
	Quantity     int64   `json:"quantity"`
}

type TransactionRequest struct {
	CampaignID int64  `json:"campaign_id"`
	UserID     string `json:"user_id"`
	Offset     int64  `json:"offset"`
	Limit      int64  `json:"limit"`
}

type TransactionResponse struct {
	Meta dtos.Meta          `json:"meta"`
	Data []*TransactionData `json:"data"`
}

type TransactionData struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	CampaignID int64   `json:"campaign_id"`
	RuleID     int64   `json:"rule_id"`
	RewardID   int64   `json:"reward_id"`
	UnitValue  int64   `json:"unit_value"`
	Quantity   int64   `json:"quantity"`
	SourceType string  `json:"source_type"`
	CreatedAt  int64   `json:"created_at"`
	Reward     *Reward `json:"reward"`
}

type GetRewardsByCampaignIDRequest struct {
	CampaignID int64   `json:"campaign_id"`
	TagIDs     []int64 `json:"tag_ids"`
	RewardName string  `json:"reward_name"`
	Cursor     string  `json:"cursor"`
	Page       int     `json:"page"`
}

func encodeGetRewardRequest(_ context.Context, req interface{}) (interface{}, error) {
	rewardID, _ := req.(int64)
	return &gami_protobuf.GetRewardRequest{
		ID: rewardID,
	}, nil
}

func encodeGetListRewardByIDsRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, _ := req.(*GetListRewardByIDsRequest)
	return &gami_protobuf.GetListRewardRequest{
		IDs: r.IDs,
	}, nil
}

func decodeGetListRewardResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetListRewardResponse)
	if r.Meta.Status != http.StatusOK {
		return &ListRewardResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	var rewardsData = make([]*Reward, len(r.Data))
	for index, item := range r.Data {
		var reward Reward
		_ = copier.Copy(&reward, item)
		rewardsData[index] = &reward
	}

	data := ListRewardResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: rewardsData,
	}
	return &data, nil
}

func decodeGetRewardResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetRewardResponse)

	if r.Meta.Status != http.StatusOK {
		return &GetRewardResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	rewardData := Reward{}
	_ = copier.Copy(&rewardData, r.Data)

	data := GetRewardResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: &rewardData,
	}

	return &data, nil
}

func encodeGetRewardPoolRequest(_ context.Context, req interface{}) (request interface{}, err error) {
	r, _ := req.(GetRewardPoolRequest)
	return &gami_protobuf.GetRewardPoolRequest{
		CampaignID: r.CampaignID,
	}, nil
}

func encodeGetRewardsByCampaignIDRequest(_ context.Context, req interface{}) (request interface{}, err error) {
	r, _ := req.(GetRewardsByCampaignIDRequest)
	return &gami_protobuf.GetRewardsByCampaignIDRequest{
		Page:       int64(r.Page),
		TagIDs:     r.TagIDs,
		Cursor:     r.Cursor,
		CampaignID: r.CampaignID,
		RewardName: r.RewardName,
	}, nil
}
func decodeGetRewardsByCampaignIDResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetRewardsByCampaignIDResponse)
	if r.Meta.Meta.Code != http.StatusOK {
		return &GetRewardsByCampaignIDResponse{
			Meta: dtos.CursorMeta{
				MetaV2: &dtos.MetaV2{
					Code:    r.Meta.Meta.Code,
					Message: r.Meta.Meta.Message,
					Cursor:  "",
				},
				Total: 0,
			},
			Data: nil,
		}, nil
	}
	var rewardData []*RewardPool
	_ = copier.Copy(&rewardData, r.Data)
	data := GetRewardsByCampaignIDResponse{
		Meta: dtos.CursorMeta{
			MetaV2: &dtos.MetaV2{
				Code:    r.Meta.Meta.Code,
				Message: r.Meta.Meta.Message,
				Cursor:  r.Meta.Meta.Cursor,
			},
			Total: r.Meta.Total,
		},
		Data: rewardData,
	}
	return &data, nil
}

func decodeGetRewardPoolResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetRewardPoolResponse)

	if r.Meta.Status != http.StatusOK {
		return &GetRewardPoolResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	var rewardData []*RewardPool
	_ = copier.Copy(&rewardData, r.Data)

	data := GetRewardPoolResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: rewardData,
	}

	return &data, nil
}

func encodeRedeemRequest(_ context.Context, req interface{}) (request interface{}, err error) {
	r, _ := req.(*RedeemRequest)
	return &gami_protobuf.RedeemRequest{
		CampaignID:   r.CampaignID,
		UserID:       r.UserID,
		Quantity:     int64(r.Quantity),
		RewardPoolID: r.RewardPoolID,
		PoolID:       r.PoolID,
		ExtraData:    r.ExtraData,
	}, nil
}

func decodeRedeemResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.RedeemResponse)
	if r == nil {
		return nil, errors.New("response is null")
	}
	if r.Meta.Status != http.StatusOK {
		resp := RedeemResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}
		if r.Meta.Data != nil {
			resp.Meta.Data = &dtos.Data{
				AmountLeft: r.Meta.Data.AmountLeft,
			}
		}
		return &resp, nil
	}

	redeemData := RedeemData{}
	_ = copier.Copy(&redeemData, r.Data)

	data := RedeemResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: &redeemData,
	}

	return &data, nil
}

func encodeTransactionRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, _ := req.(*TransactionRequest)
	return &gami_protobuf.GetTransactionRequest{
		CampaignID: r.CampaignID,
		UserID:     r.UserID,
		Limit:      r.Limit,
		Offset:     r.Offset,
	}, nil
}

func decodeTransactionResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetTransactionResponse)
	if r.Meta.Status != http.StatusOK {
		return &TransactionResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	var redeemData []*TransactionData
	_ = copier.Copy(&redeemData, r.Data)

	data := TransactionResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: redeemData,
	}
	return &data, nil
}
