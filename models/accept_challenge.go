package models

import "time"

// AcceptChallenge models
type AcceptChallenge struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	CampaignID int64     `json:"campaign_id"`
	CreatedAt  time.Time `json:"created_at"`
}

// TableName accept_challenge
func (AcceptChallenge) TableName() string {
	return "accept_challenge"
}
