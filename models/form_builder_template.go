package models

type FormBuilderTemplateStatus string

const (
	FormBuilderStatusActive   FormBuilderTemplateStatus = "ACTIVE"
	FormBuilderStatusInactive FormBuilderTemplateStatus = "INACTIVE"
)

func (cs FormBuilderTemplateStatus) String() string {
	return string(cs)
}

type FormBuilderTemplate struct {
	Base
	TenantID    int64                     `json:"tenant_id"`
	Code        string                    `json:"code"`
	Name        string                    `json:"name"`
	Type        string                    `json:"type"`
	Description string                    `json:"description"`
	Status      FormBuilderTemplateStatus `json:"status"`
	Template    string                    `json:"template"`
}
