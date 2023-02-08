package models

import (
	"time"

	"github.com/google/uuid"
)

type ScoreTransaction struct {
	ID         uuid.UUID  `gorm:"primary_key;"`
	CampaignID int64      `json:"campaign_id"`
	UserID     int64      `json:"user_id"`
	Amount     int64      `json:"amount"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
}

func (ScoreTransaction) TableName() string {
	return "score_transactions"
}
