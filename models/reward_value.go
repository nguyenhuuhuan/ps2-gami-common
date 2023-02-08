package models

// RewardValue model
type RewardValue struct {
	Base
	CampaignID int64  `gorm:"column:campaign_id"`
	RewardID   int64  `gorm:"column:reward_id"`
	UserID     string `gorm:"column:user_id"`
	Amount     int    `gorm:"column:amount"`
}

// TableName define
func (RewardValue) TableName() string {
	return "reward_values"
}
