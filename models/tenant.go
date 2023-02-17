package models

type TenantStatus string

const (
	TenantStatusActive   TenantStatus = "ACTIVE"
	TenantStatusInactive TenantStatus = "INACTIVE"
)

type Tenant struct {
	Base
	Code   string       `json:"code"`
	Name   string       `json:"name"`
	Status TenantStatus `json:"status"`
}

func (Tenant) TableName() string {
	return "tenants"
}
