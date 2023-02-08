package models

// ForcedUpdate model
type ForcedUpdate struct {
	Base
	Platform string
	Version  string
	Module   string
}
