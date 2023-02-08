package models

// CampaignValue model
type CampaignValue struct {
	Base
	CampaignID int64 `json:"campaign_id"`
	UserID     int64 `json:"user_id"`
	Amount     int   `json:"amount"`
	RuleID     int64 `json:"rule_id"`
}

// TableName define
func (CampaignValue) TableName() string {
	return "campaign_values"
}
