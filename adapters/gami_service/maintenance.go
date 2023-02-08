package gami_service

import (
	"context"
	"errors"
	"net/http"

	"gitlab.id.vin/gami/ps2-gami-common/dtos"
	gamiProtobuf "gitlab.id.vin/gami/ps2-gami-proto/pb"
)

type MaintenanceStatusCampaignRequest struct {
	CampaignType string `json:"campaign_type"`
}

type MaintenanceStatusCampaignResponse struct {
	Meta dtos.Meta `json:"meta"`
}

func encodeMaintenanceStatusCampaignRequest(_ context.Context, req interface{}) (interface{}, error) {
	r, ok := req.(*MaintenanceStatusCampaignRequest)
	if !ok {
		return nil, errors.New("encode - maintenance status for campaign, cast error")
	}

	return &gamiProtobuf.MaintenanceStatusCampaignRequest{
		CampaignType: r.CampaignType,
	}, nil
}

func decodeMaintenanceStatusCampaignResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, ok := resp.(*gamiProtobuf.MaintenanceStatusCampaignResponse)
	if !ok {
		return nil, errors.New("decode - maintenance status for campaign, cast error")
	}

	if r.Meta.Status != http.StatusOK {
		return &MaintenanceStatusCampaignResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	return &MaintenanceStatusCampaignResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
	}, nil
}
