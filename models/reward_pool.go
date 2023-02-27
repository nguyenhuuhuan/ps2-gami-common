package models

import (
	"gorm.io/gorm"
)

// RewardPoolStatus Status define
type RewardPoolStatus string

// RewardPool const
const (
	RewardPoolActive   RewardPoolStatus = "ACTIVE"
	RewardPoolInActive RewardPoolStatus = "INACTIVE"
)

// RewardPool model
type RewardPool struct {
	Base
	PoolID        int64            `gorm:"column:pool_id" json:"pool_id"`
	Pool          *Pool            `json:"pool,omitempty"`
	RewardID      int64            `gorm:"column:reward_id" json:"reward_id"`
	Reward        *Reward          `json:"reward,omitempty"`
	Rate          float64          `gorm:"column:rate" json:"rate"`
	Amount        int              `gorm:"column:amount" json:"amount"`
	AmountLeft    int              `gorm:"column:amount_left" json:"amount_left"`
	AlertAmount   int              `gorm:"column:alert_amount" json:"alert_amount"`
	UnitValue     int              `gorm:"column:unit_value" json:"unit_value"`
	Status        RewardPoolStatus `gorm:"column:status" json:"status"`
	Probability   float64          `gorm:"-"`
	ForCampaignID int64            `gorm:"-"`
	TenantID      int64            `json:"tenant_id"`
}

// BeforeCreate Reward Pool
func (p *RewardPool) BeforeCreate(tx *gorm.DB) error {
	p.AmountLeft = p.Amount
	return nil
}

// TableName define
func (RewardPool) TableName() string {
	return "reward_pools"
}
