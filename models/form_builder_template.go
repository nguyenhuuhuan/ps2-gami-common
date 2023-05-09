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
	BaseWithoutDeletedAt
	TenantID    int64                     `json:"tenant_id"`
	Name        string                    `json:"name"`
	Type        string                    `json:"type"`
	Description string                    `json:"description"`
	Status      FormBuilderTemplateStatus `json:"status"`
	Template    JSON                      `json:"template"`
}
