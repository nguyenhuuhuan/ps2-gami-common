package models

type RuleEventTemplate struct {
	BaseWithoutDeletedAt
	Name        string `json:"name"`
	Code        string `gorm:"<-:create" json:"code"`
	Payload     *JSON  `json:"payload,omitempty"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	Event       string `json:"event"`
}

// TableName rule_event_templates
func (RuleEventTemplate) TableName() string {
	return "rule_event_templates"
}
