package models

// TemplateStatus is a string type
type TemplateStatus string

// TemplateType is a string type
type TemplateType string

// Template status and Template type enum definition
const (
	TemplateStatusActive   TemplateStatus = "ACTIVE"
	TemplateStatusInactive TemplateStatus = "INACTIVE"
	TemplateGenericType    TemplateType   = "GENERIC"
	TemplateSpecifiedType  TemplateType   = "SPECIFIED"
	TemplateCampaignType   TemplateType   = "CAMPAIGN"
	TemplateManualCodeType TemplateType   = "MANUAL_CODE"
)

const (
	BillPayFirstTransaction = "BILL_001"
)

// ConditionTemplate models
type ConditionTemplate struct {
	Base
	Name        string
	Code        string
	Description string
	Status      TemplateStatus
	Statement   string
	AttrIDs     string
	Type        TemplateType
}
