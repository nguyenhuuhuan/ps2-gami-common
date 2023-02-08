package models

type MaintenanceType string

const (
	MaintenanceTypeScreen   MaintenanceType = "SCREEN"
	MaintenanceTypeCampaign MaintenanceType = "CAMPAIGN"
)

// String returns MaintenanceType as a string.
func (cs MaintenanceType) String() string {
	return string(cs)
}

type Maintenance struct {
	BaseWithoutDeletedAt
	Type         MaintenanceType `json:"type"`
	Status       bool            `json:"status"`
	CampaignType CampaignType    `json:"campaign_type"`
	ScreenCodeID int64           `json:"screen_code_id"`
	ScreenCode   *ScreenCode     `json:"screen_code"`
}

// TableName maintenances
func (Maintenance) TableName() string {
	return "maintenances"
}
