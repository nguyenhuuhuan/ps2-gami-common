package models

import "time"

type UserScore struct {
	CampaignID int64      `json:"campaign_id" gorm:"primary_key"`
	UserID     int64      `json:"user_id" gorm:"primary_key"`
	TotalScore int64      `json:"total_score"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"`
}

// TableName returns table name
func (UserScore) TableName() string {
	return "user_scores"
}
