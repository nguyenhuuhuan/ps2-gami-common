package models

import "time"

// BlackWhiteListType is a string type.
type BlackWhiteListType string

// Group status enum definition
const (
	BlackListType BlackWhiteListType = "BLACKLIST"
	WhiteListType BlackWhiteListType = "WHITELIST"
)

// BlackWhiteList models
type BlackWhiteList struct {
	ID         int64              `json:"id"`
	CampaignID int64              `json:"campaign_id"`
	GroupID    int64              `json:"group_id"`
	Type       BlackWhiteListType `json:"type"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

// TableName black_white_list
func (BlackWhiteList) TableName() string {
	return "black_white_list"
}
