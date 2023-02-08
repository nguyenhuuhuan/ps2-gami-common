package models

type WithdrawMiddleReward struct {
	BaseWithoutDeletedAt
	CampaignID int64  `gorm:"column:campaign_id" json:"campaign_id"`
	UserID     string `gorm:"column:user_id" json:"user_id"`
	Amount     int64  `gorm:"column:amount" json:"amount"`
	Reason     string `gorm:"column:reason" json:"reason"`
	Maker      string `gorm:"column:maker" json:"maker"`
	Checker    string `gorm:"column:checker" json:"checker"`
}

// TableName withdraw_middle_rewards
func (WithdrawMiddleReward) TableName() string {
	return "withdraw_middle_rewards"
}
