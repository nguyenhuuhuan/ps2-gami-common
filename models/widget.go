package models

import "time"

// WidgetStatus is a string type.
type WidgetStatus string

// WidgetRuleRequirement is a string type
type WidgetRuleRequirement string

// String returns WidgetStatus as a string.
func (cs WidgetStatus) String() string {
	return string(cs)
}

// Widget status enum definition
const (
	WidgetStatusActive   WidgetStatus = "ACTIVE"
	WidgetStatusInactive WidgetStatus = "INACTIVE"
)

// Widget models
type Widget struct {
	Base
	HomeVinID            bool         `json:"home_vin_id"`
	HomePromotion        bool         `json:"home_promotion"`
	LoyaltyProgramScreen bool         `json:"loyalty_program_screen"`
	DisplayOrder         int          `json:"display_order"`
	Title                string       `json:"title"`
	TitleColor           string       `json:"title_color"`
	Status               WidgetStatus `json:"status"`
	Icon                 string       `json:"icon"`
	Banner               string       `json:"banner"`
	Lottie               string       `json:"lottie"`
	AppType              string       `json:"app_type"`
	DeepLink             string       `json:"deep_link"`
	CTATitle             string       `json:"cta_title"`
	CTAColor             string       `json:"cta_color"`
	LoyaltyProgram       string       `json:"loyalty_program"`
	ExpiryCTATitle       string       `json:"expiry_cta_title"`
	ExpiryDeepLink       string       `json:"expiry_deep_link"`
	EffectiveCTATitle    string       `json:"effective_cta_title"`
	EffectiveDeepLink    string       `json:"effective_deep_link"`
	DateStart            *time.Time   `json:"date_start"`
	DateEnd              *time.Time   `json:"date_end"`
	EffectiveFromDate    *time.Time   `json:"effective_from_date"`
	EffectiveToDate      *time.Time   `json:"effective_to_date"`
	Campaign             *Campaign    `json:"campaign,omitempty"`
	WidgetGroup          *WidgetGroup `json:"widget_group"`
	Players              int64        `gorm:"-" json:"players"`
	CampaignID           int64        `json:"campaign_id"`
	Config               *JSON        `sql:"type:json" json:"config,omitempty"`
	IsShowProgress       bool         `json:"is_show_progress"`
}

// DateStartStr to convert to string when get detail
func (m *Widget) DateStartStr() string {
	return m.DateStart.Format(DateFormat)
}

// DateEndStr function to convert to string when get detail
func (m *Widget) DateEndStr() string {
	return m.DateEnd.Format(DateFormat)
}
