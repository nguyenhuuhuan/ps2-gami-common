package models

// AttributeValue models
type AttributeValue struct {
	Base
	AttributeID int64      `gorm:"column:attribute_id" json:"attribute_id"`
	Attribute   *Attribute `gorm:"association_autoupdate:false" json:"attribute"`
	ConditionID int64      `gorm:"column:condition_id" json:"condition_id"`
	Value       string     `gorm:"column:value" json:"value"`
}
