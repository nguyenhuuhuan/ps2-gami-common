package dtos

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type Auth0Claims struct {
	jwt.StandardClaims
	VinIdInfo VinIdInfo `json:"https://vinid.net"`
}
type VinIdInfo struct {
	UserID      int64  `json:"user_id"`
	Level       int64  `json:"level"`
	PhoneNumber string `json:"phone_number"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5t string   `json:"x5t"`
	X5c []string `json:"x5c"`
}
