package gami_service

import (
	"context"
	"github.com/jinzhu/copier"
	"gitlab.id.vin/gami/ps2-gami-common/dtos"
	gami_protobuf "gitlab.id.vin/gami/ps2-gami-proto/pb"
	"net/http"
)

type BlackWhiteList struct {
	CampaignID int64  `json:"campaign_id"`
	GroupID    int64  `json:"group_id"`
	Type       string `json:"type"`
}
type GetBlackWhiteListResponse struct {
	Meta dtos.Meta       `json:"meta"`
	Data *BlackWhiteList `json:"data"`
}

func encodeGetBlackWhiteListRequest(_ context.Context, req interface{}) (interface{}, error) {
	campaignID, _ := req.(int64)
	return &gami_protobuf.GetBlackWhiteListRequest{
		CampaignID: campaignID,
	}, nil
}

func decodeGetBlackWhiteListResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetBlackWhiteListResponse)
	if r.Meta.Status != http.StatusOK {
		return &GetBlackWhiteListResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	var blackWhiteListData BlackWhiteList
	_ = copier.Copy(&blackWhiteListData, r.Data)

	data := GetBlackWhiteListResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: &blackWhiteListData,
	}
	return &data, nil
}
