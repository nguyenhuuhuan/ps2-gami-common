package models

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

// Transaction model
type Transaction struct {
	ID         string     `json:"id" gorm:"type:uuid;primaryKey;"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"`
	UserID     string     `json:"user_id"`
	RewardID   int64      `json:"reward_id"`
	Reward     *Reward    `json:"reward,omitempty"`
	CampaignID int64      `json:"campaign_id"`
	Campaign   *Campaign  `json:"campaign,omitempty"`
	RuleID     *int64     `json:"rule_id"`
	UnitValue  int        `json:"unit_value"`
	SourceType string     `gorm:"column:source_type" json:"source_type"`
	Quantity   int        `json:"quantity"`
	ExtraData  *JSON      `sql:"type:json" json:"extra_data,omitempty"`
	PoolID     *int64     `json:"pool_id"`
}

// BeforeCreate transaction uuid
func (m *Transaction) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = xid.New().String()
	}
	return nil
}

func (m *Transaction) Marshal() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(m)
}

func (m *Transaction) Unmarshal(data []byte) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, m)
}
