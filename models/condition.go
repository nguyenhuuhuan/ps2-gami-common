package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// ConditionStatus is a string type
type ConditionStatus string

// Condition Status enum definition
const (
	ConditionStatusActive   ConditionStatus = "ACTIVE"
	ConditionStatusInactive ConditionStatus = "INACTIVE"
)

// Condition models
type Condition struct {
	Base
	PoolID              int64              `json:"pool_id"`
	Pool                *Pool              `json:"pools,omitempty"`
	ConditionTemplateID int64              `gorm:"column:template_id" json:"condition_template_id"`
	ConditionTemplate   *ConditionTemplate `json:"condition_templates,omitempty"`
	RewardID            int64              `json:"reward_id"`
	Reward              *Reward            `json:"rewards,omitempty"`
	Name                string             `json:"name"`
	AttributeValue      []*AttributeValue  `json:"attribute_values,omitempty"`
	Status              ConditionStatus    `gorm:"column:status" json:"status"`
	IncludeReward       *IncludeReward     `sql:"type:include_reward" json:"include_reward,omitempty"`
	ExcludeReward       *ExcludeReward     `sql:"type:exclude_reward" json:"exclude_reward,omitempty"`
}

type ConditionResult struct {
	Result bool `json:"result"`
}

type IncludeReward []int64
type ExcludeReward []int64

// Value function call when save a row mystery box campaign into db
func (c IncludeReward) Value() (driver.Value, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

// Function to convert to slice
func (c *IncludeReward) ToSlice() []int64 {
	var includeIds []int64
	for _, value := range *c {
		includeIds = append(includeIds, value)
	}
	return includeIds
}

// Scan function call when query from db
func (c *IncludeReward) Scan(value interface{}) error {
	if value == nil {
		*c = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source for include reward")
	}
	err := json.Unmarshal(s, c)
	return err
}

// Value function call when save a row mystery box campaign into db
func (c ExcludeReward) Value() (driver.Value, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

// Value function call when save a row mystery box campaign into db
func (c *ExcludeReward) Scan(value interface{}) error {
	if value == nil {
		*c = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source for exclude reward")
	}
	err := json.Unmarshal(s, c)
	return err
}

// Function to convert to slice
func (c *ExcludeReward) ToSlice() []int64 {
	var includeIds []int64
	for _, value := range *c {
		includeIds = append(includeIds, value)
	}
	return includeIds
}
