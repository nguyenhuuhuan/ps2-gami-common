package models

type VisualizationTag struct {
	BaseWithoutDeletedAt
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
