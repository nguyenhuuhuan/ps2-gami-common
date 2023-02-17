package models

type Tenant struct {
	Base
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}
