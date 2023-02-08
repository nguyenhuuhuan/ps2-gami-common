package models

// WidgetGroup models
type WidgetGroup struct {
	Base
	WidgetID int64              `json:"widget_id"`
	GroupID  int64              `json:"group_id"`
	Type     BlackWhiteListType `json:"type"`
}

// TableName widget_group
func (WidgetGroup) TableName() string {
	return "widget_group"
}
