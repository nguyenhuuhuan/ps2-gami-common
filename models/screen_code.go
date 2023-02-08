package models

type ScreenCode struct {
	BaseWithoutDeletedAt
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	Description string       `json:"description"`
	Maintenance *Maintenance `json:"maintenance"`
}

// TableName rule_event_templates
func (ScreenCode) TableName() string {
	return "screen_codes"
}
