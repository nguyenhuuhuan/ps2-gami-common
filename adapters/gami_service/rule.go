package gami_service

import (
	"context"
	"net/http"

	"github.com/jinzhu/copier"

	gami_protobuf "gitlab.id.vin/gami/ps2-gami-proto/pb"

	"gitlab.id.vin/gami/ps2-gami-common/dtos"
)

// GetRuleRequest return campaign rules
type GetRuleRequest struct {
	CampaignID int64 `json:"campaign_id"`
}

// GetRulesResponse return campaign rules
type GetRulesResponse struct {
	Meta dtos.Meta   `json:"meta"`
	Data []*RuleItem `json:"data"`
}

// RuleItem struct
type RuleItem struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	RuleID    int64  `json:"rule_id"`
	Quantity  int64  `json:"quantity"`
	Icon      string `json:"icon"`
	Deeplink  string `json:"deeplink"`
	IsShown   bool   `json:"is_shown"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func encodeGetRuleRequest(_ context.Context, req interface{}) (interface{}, error) {
	campaignID, _ := req.(int64)
	return &gami_protobuf.GetRuleRequest{
		CampaignID: campaignID,
	}, nil
}

func decodeGetRuleResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r, _ := resp.(*gami_protobuf.GetRuleResponse)
	if r.Meta.Status != http.StatusOK {
		return &TransactionResponse{
			Meta: dtos.Meta{
				Code:    int(r.Meta.Status),
				Message: r.Meta.Message,
			},
		}, nil
	}

	var ruleData []*RuleItem
	_ = copier.Copy(&ruleData, r.Data)

	data := GetRulesResponse{
		Meta: dtos.Meta{
			Code:    int(r.Meta.Status),
			Message: r.Meta.Message,
		},
		Data: ruleData,
	}
	return &data, nil
}
