package models

import (
	"time"
)

// Base Models
type Base struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}

// BaseWithoutDeletedAt Models
type BaseWithoutDeletedAt struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreatedTimestamp Convert to timestamp
func (b *Base) CreatedTimestamp() int64 {
	return b.CreatedAt.Unix()
}

// CreatedTimestamp Convert to timestamp
func (b *Base) UpdatedTimestamp() int64 {
	return b.UpdatedAt.Unix()
}
