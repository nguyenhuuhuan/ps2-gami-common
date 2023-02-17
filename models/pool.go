package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// PoolStatus define
type PoolStatus string

// const Pool
const (
	PoolStatusActive   PoolStatus = "ACTIVE"
	PoolStatusInactive PoolStatus = "INACTIVE"
)

// String returns PoolStatus as a string.
func (cs PoolStatus) String() string {
	return string(cs)
}

// Pool model
type Pool struct {
	Base
	TenantID   int64         `json:"tenant_id"`
	Campaign   *Campaign     `json:"campaign,omitempty"`
	CampaignID *int64        `json:"campaign_id"`
	RewardPool []*RewardPool `json:"reward_pools,omitempty"`
	Condition  []*Condition  `json:"conditions,omitempty"`
	Status     PoolStatus    `json:"status"`
	Code       string        `gorm:"<-:create" json:"code"`
	Maker      string        `json:"maker"`
	Checker    string        `json:"checker"`
	Priority   int64         `json:"priority"`
	StartDate  *time.Time    `json:"start_date,omitempty"`
	EndDate    *time.Time    `json:"end_date,omitempty"`
}

// AfterDelete func
func (pool *Pool) AfterDelete(tx *gorm.DB) error {
	return tx.Delete(&RewardPool{}, "pool_id = ?", pool.ID).Error
}

func (pool *Pool) ExistGemReward() bool {
	if len(pool.RewardPool) > 0 {
		for _, _item := range pool.RewardPool {
			if _item != nil && _item.Reward != nil && _item.Reward.Type == RewardTypeGem {
				return true
			}
		}
	}

	return false
}
