package models

import "encoding/json"

// Redeem Message
type Redeem struct {
	Source       string          `json:"source"`
	RuleID       int64           `json:"rule_id"`
	RuleCode     string          `json:"rule_code"`
	UserID       string          `json:"user_id"`
	TenantID     int64           `json:"tenant_id"`
	TenantCode   string          `json:"tenant_code"`
	Name         string          `json:"name"`
	Code         string          `json:"code"`
	IsManual     bool            `json:"is_manual"`
	RuleMasterID int64           `json:"rule_master_id,omitempty"`
	CampaignID   int64           `json:"campaign_id"`
	CampaignCode string          `json:"campaign_code"`
	Quantity     int             `json:"quantity"`
	ExtraData    json.RawMessage `json:"addition"`
}

// AdditionData define addition data from PE
type AdditionData struct {
	EventType    string      `json:"event_type"`
	EventPayload interface{} `json:"event_payload"`
}

// LotteryMessage define message for kafka
type LotteryMessage struct {
	CampaignID int64  `json:"campaign_id"`
	UserID     string `json:"user_id"`
}

// MysteryBoxMessage define message for kafka
type MysteryBoxMessage struct {
	CampaignCode string `json:"campaign_code"`
	UserID       string `json:"user_id"`
	RuleName     string `json:"rule_name"`
	Quantity     int64  `json:"quantity"`
}

// GemMessage define message for gem
type GemMessage struct {
	CampaignCode string `json:"campaign_code"`
	Quantity     int    `json:"quantity"`
	UserID       string `json:"user_id"`
}

// RewardForwardMessage define message to forward to other campaign
type RewardForwardMessage struct {
	Quantity       int    `json:"quantity"`
	UserID         string `json:"user_id"`
	RuleID         *int64 `json:"rule_id,omitempty"`
	FromCampaignID int64  `json:"from_campaign_id"`
	ToCampaignID   int64  `json:"to_campaign_id"`
}

// VoucherMessage define message for voucher
type VoucherMessage struct {
	VoucherCode string `json:"voucher_code"`
	Quantity    int    `json:"num_serial"`
	UserID      int64  `json:"user_id"`
}

// CollectionMessage define message for kafka
type CollectionMessage struct {
	CampaignCode string `json:"campaign_code"`
	UserID       string `json:"user_id"`
	RuleID       int64  `json:"rule_id"`
	Quantity     int64  `json:"quantity"`
}

type InitRewardPoolQuantityMessage struct {
	CampaignID   int64 `json:"campaign_id"`
	RewardID     int64 `json:"reward_id"`
	RewardPoolID int64 `json:"reward_pool_id"`
}

type UpdateRewardPoolQuantityMessage struct {
	UnitValue    int   `json:"unit_value"`
	Quantity     int   `json:"quantity"`
	RewardPoolID int64 `json:"reward_pool_id"`
}

type ItemSetRedeemMessage struct {
	ID          string               `json:"id"`
	RequestID   string               `json:"request_id"`
	RefEventID  string               `json:"ref_event_id"`
	Event       string               `json:"event"`
	ServiceCode string               `json:"service_code"`
	Timestamp   int64                `json:"timestamp"`
	UserID      string               `json:"user_id"`
	PayloadID   string               `json:"payload_id"`
	Payload     ItemSetRedeemPayload `json:"payload"`
}

type ItemSetRedeemPayload struct {
	FromCampaignID int64  `json:"from_campaign_id"`
	ToCampaignID   int64  `json:"to_campaign_id"`
	UserID         string `json:"user_id"`
	RewardPoolID   int64  `json:"reward_pool_id"`
	RewardID       int64  `json:"reward_id"`
	Amount         int    `json:"amount"`
}

type QuestPassRuleMessage struct {
	ID          string               `json:"id"`
	RequestID   string               `json:"request_id"`
	RefEventID  string               `json:"ref_event_id"`
	Event       string               `json:"event"`
	ServiceCode string               `json:"service_code"`
	Timestamp   int64                `json:"timestamp"`
	UserID      string               `json:"user_id"`
	PayloadID   string               `json:"payload_id"`
	Payload     QuestPassRulePayload `json:"payload"`
}

type QuestPassRulePayload struct {
	CampaignID int64  `json:"campaign_id"`
	UserID     string `json:"user_id"`
	RuleID     int64  `json:"rule_id"`
	Quantity   int64  `json:"quantity"`
}
