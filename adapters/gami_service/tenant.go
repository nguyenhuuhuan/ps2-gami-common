package gami_service

import (
	"context"
	"net/http"

	"github.com/jinzhu/copier"
	"gitlab.id.vin/gami/ps2-gami-common/dtos"
	gami_protobuf "gitlab.id.vin/gami/ps2-gami-proto/pb"
)

type GetTenantRequest struct {
	UserID int64 `json:"user_id"`
}

type GetTenantResponse struct {
	Meta dtos.Meta `json:"meta"`
	Data *Tenant   `json:"data"`
}

type Tenant struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func encodeGetTenantRequest(_ context.Context, req interface{}) (interface{}, error) {
	userID, _ := req.(int64)
	return &gami_protobuf.GetTenantRequest{
		UserId: userID,
	}, nil
}

func decodeGetTenantResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetTenantResponse)

	if r.Meta.Status != http.StatusOK {
		return &GetTenantResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	tenant := Tenant{}
	_ = copier.Copy(&tenant, r.Data)

	data := GetTenantResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: &tenant,
	}

	return &data, nil
}
