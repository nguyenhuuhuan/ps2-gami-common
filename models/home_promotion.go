package models

type HomePromotionStatus string

const (
	HomePromotionStatusActive   HomePromotionStatus = "ACTIVE"
	HomePromotionStatusInactive HomePromotionStatus = "INACTIVE"
)

type HomePromotion struct {
	Base
	Name        string              `json:"name"`
	Code        string              `gorm:"<-:create" json:"code"`
	Status      HomePromotionStatus `json:"status"`
	Description string              `json:"description"`
	Config      *JSON               `sql:"type:json" json:"config,omitempty"`
	Component   []*Component        `json:"component,omitempty"`
}

// TableName home_promotions
func (HomePromotion) TableName() string {
	return "home_promotions"
}

// String returns HomePromotionStatus as a string.
func (cs HomePromotionStatus) String() string {
	return string(cs)
}
