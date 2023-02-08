package models

type ComponentType string
type ComponentStatus string
type ComponentChipType string

const (
	ComponentFaq                     ComponentType = "FAQ"
	ComponentChips                   ComponentType = "CHIPS"
	ComponentReferral                ComponentType = "REFERRAL"
	ComponentHighlight               ComponentType = "HIGHLIGHT_ITEM"
	ComponentDailyCheckIn            ComponentType = "DAILY_CHECK_IN"
	ComponentCampaignNormal          ComponentType = "CAMPAIGN_NORMAL"
	ComponentDailyCheckInProgression ComponentType = "DAILY_CHECK_IN_PROGRESSION"
	ComponentVoucher                 ComponentType = "VOUCHER"
	ComponentSpecialBanner           ComponentType = "SPECIAL_BANNER"
	ComponentHomeUserProfile         ComponentType = "HOME_PROFILE"
	ComponentCampaignMap             ComponentType = "CAMPAIGN_MAP"
	ComponentSlider                  ComponentType = "SLIDER"
	ComponentLandingPage             ComponentType = "LANDING_PAGE"

	// Component status
	ComponentStatusActive   ComponentStatus = "ACTIVE"
	ComponentStatusInactive ComponentStatus = "INACTIVE"
	// Chip type
	ChipComponentRewardType ComponentChipType = "REWARD"
	ChipComponentNormalType ComponentChipType = "NORMAL"
)

type Component struct {
	Base
	Name            string          `json:"name"`
	Type            ComponentType   `gorm:"<-:create" json:"type"`
	Status          ComponentStatus `json:"status"`
	ProgramCode     string          `json:"program_code"`
	Config          *JSON           `sql:"type:json" json:"config,omitempty"`
	HomePromotion   HomePromotion   `json:"home_promotion"`
	HomePromotionID int64           `json:"home_promotion_id"`
}

// TableName components
func (Component) TableName() string {
	return "components"
}

// String returns ComponentType as a string.
func (cs ComponentType) String() string {
	return string(cs)
}

// String returns ComponentStatus as a string.
func (cs ComponentStatus) String() string {
	return string(cs)
}

// String returns ComponentChipType as a string.
func (cs ComponentChipType) String() string {
	return string(cs)
}
