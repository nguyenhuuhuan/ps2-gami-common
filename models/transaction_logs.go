package models

import (
	"time"

	"github.com/kjk/betterguid"
	"gorm.io/gorm"
)

type TransactionLogs struct {
	ID                string    `gorm:"type:uuid;primaryKey;"`
	UserID            string    `json:"user_id"`
	TenantID          int64     `json:"tenant_id"`
	CampaignID        int64     `json:"campaign_id"`
	Campaign          *Campaign `json:"campaign,omitempty"`
	RewardID          int64     `json:"reward_id"`
	Reward            *Reward   `json:"reward,omitempty"`
	SourceType        string    `gorm:"column:source_type"`
	TransactionRef    *string   `gorm:"column:transaction_ref" json:"transaction_ref"`
	RawTransactionRef *string   `gorm:"column:raw_transaction_ref" json:"raw_transaction_ref"`
	Request           *JSON     `sql:"type:json" json:"request,omitempty"`
	Response          *JSON     `sql:"type:json" json:"response,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	IsSuccess         bool      `json:"is_success"`
}

// BeforeCreate transaction uuid
func (m *TransactionLogs) BeforeCreate(tx *gorm.DB) error {
	m.ID = betterguid.New()
	return nil
}
