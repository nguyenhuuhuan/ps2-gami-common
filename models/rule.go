package models

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

// RuleStatus is a string type.
type RuleStatus string

// String returns MerchantStatus as a string.
func (cs RuleStatus) String() string {
	return string(cs)
}

// IsActive convert from string to bool
func (cs RuleStatus) IsActive() bool {
	return cs == RuleStatusACTIVE
}

// Rule Status Define
const (
	RuleStatusACTIVE   RuleStatus = "ACTIVE"
	RuleStatusINACTIVE RuleStatus = "INACTIVE"
)

// Rule Models
type Rule struct {
	Base
	TenantID     int64      `json:"tenant_id"`
	Name         string     `json:"name"`
	RuleID       *int64     `json:"rule_id"`
	Campaign     *Campaign  `json:"campaign,omitempty"`
	CampaignID   int64      `json:"campaign_id"`
	IsManual     bool       `json:"is_manual"`
	Content      *string    `json:"content"`
	Icon         string     `json:"icon"`
	Status       RuleStatus `json:"status"`
	Quantity     int64      `json:"quantity"`
	Deeplink     string     `json:"deeplink"`
	RewardPoolID int64      `gorm:"-" json:"reward_pool_id,omitempty"`
	PoolID       int64      `gorm:"-" json:"pool_id,omitempty"`
	PoolCode     string     `gorm:"-" json:"pool_code,omitempty"`
	RewardCode   string     `gorm:"-" json:"reward_code,omitempty"`
	IsShown      bool       `json:"is_shown"`
	UniqueKey    string     `json:"unique_key,omitempty" gorm:"-" `
	ExtraData    *JSON      `json:"extra_data,omitempty" sql:"type:json"`
}

// RuleLog model
type RuleLog struct {
	ID         int64     `json:"id" gorm:"primaryKey" `
	RuleID     int64     `json:"rule_id"`
	CampaignID int64     `json:"campaign_id"`
	Campaign   *Campaign `json:"campaign,omitempty"`
	Rule       Rule      `json:"rule"`
	UserID     string    `json:"user_id"`
	TenantID   int64     `json:"tenant_id"`
	ExtraData  *JSON     `json:"extra_data,omitempty" sql:"type:json"`
	CreatedAt  time.Time `json:"created_at"`
}

func (m *RuleLog) Marshal() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(m)
}

func (m *RuleLog) Unmarshal(data []byte) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, m)
}
