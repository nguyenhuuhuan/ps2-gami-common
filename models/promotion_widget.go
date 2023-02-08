package models

import (
	"time"
)

// PromotionWidgets promotion widget orm.
type PromotionWidgets struct {
	Base
	IsHighlight       bool      `json:"is_highlight"`
	HighlightPosition int64     `json:"highlight_position"`
	HighlightBanner   string    `json:"highlight_banner"`
	BasicBanner       string    `json:"basic_banner"`
	CampaignName      string    `json:"campaign_name"`
	ShowTime          time.Time `json:"show_time"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	HideTime          time.Time `json:"hide_time"`
	Active            bool      `json:"active"`
	MetaData          *JSON     `json:"meta_data"`
}

func (PromotionWidgets) TableName() string {
	return "promotion_widgets"
}
