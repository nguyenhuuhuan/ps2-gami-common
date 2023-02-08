package models

type Tag struct {
	BaseWithoutDeletedAt
	Name string `json:"name"`
}
