package models

// AttributeType is a string type
type AttributeType string

// Attribute types: add when needed more types
const (
	AttributeTypeString   AttributeType = "string"
	AttributeTypeInteger  AttributeType = "integer"
	AttributeTypeDateTime AttributeType = "datetime"
)

// Attribute pattern: add when needed more pattern
const (
	AttributePatternLimit         string = "LIMIT_VALUE"
	AttributePatternTimeLimitFrom string = "TIME_LIMIT_FROM"
	AttributePatternTimeLimitTo   string = "TIME_LIMIT_TO"
)

// Attribute models
type Attribute struct {
	Base
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        AttributeType `json:"type"`
	Pattern     string        `json:"pattern"`
}
