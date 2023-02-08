package models

import (
	"time"
)

type JobStatus string

const (
	JobStatusActive   JobStatus = "ACTIVE"
	JobStatusInActive JobStatus = "INACTIVE"
)

type JobFrequency string

const (
	DAILY  JobFrequency = "DAILY"
	WEEKLY JobFrequency = "WEEKLY"
)

type JobType string

const (
	REWARD_RESTOCK JobType = "REWARD_RESTOCK"
)

type Job struct {
	BaseWithoutDeletedAt
	Type      JobType      `gorm:"type" json:"type"`
	FromDate  time.Time    `gorm:"from_date" json:"from_date"`
	ToDate    time.Time    `gorm:"to_date" json:"to_date"`
	Times     int64        `gorm:"times" json:"times"`
	Config    *JSON        `gorm:"config" json:"config"`
	Frequency JobFrequency `gorm:"frequency" json:"frequency"`
	Status    JobStatus    `gorm:"status" json:"status"`
}

func (j *Job) FromDateStr() string {
	return j.FromDate.Format(DateFormat)
}

func (j *Job) ToDateStr() string {
	return j.ToDate.Format(DateFormat)
}

func (js JobType) String() string {
	return string(js)
}

func (js JobFrequency) String() string {
	return string(js)
}

func (js JobStatus) String() string {
	return string(js)
}
