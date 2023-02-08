package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	ConfigRedeemKey string = "redeem_maintenance"
)

// Config models
type Config struct {
	ID        int64       `json:"id"`
	MyKey     string      `json:"key" orm:"column(my_key)"`
	MyValue   *ConfigItem `sql:"type:json" json:"value,omitempty" orm:"column(my_value)"`
	CreatedAt time.Time   `json:"created_at"`
}

// ConfigureItem struct
type ConfigItem struct {
	ItemValue interface{} `json:"value"`
}

// Value function call when save a row mystery box campaign into db
func (c ConfigItem) Value() (driver.Value, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

// Scan function call when query from db
func (c *ConfigItem) Scan(value interface{}) error {
	if value == nil {
		//*c = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source for include reward")
	}
	err := json.Unmarshal(s, c)
	return err
}

func (c *Config) GetString() (string, bool) {
	s, ok := c.MyValue.ItemValue.(string)
	return s, ok
}

func (c *Config) GetBool() (bool, bool) {
	isBool, ok := c.MyValue.ItemValue.(bool)
	return isBool, ok
}

func (c *Config) GetInt64() (int64, bool) {
	number, ok := c.MyValue.ItemValue.(int64)
	return number, ok
}
