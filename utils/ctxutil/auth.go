package ctxutil

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"gitlab.id.vin/gami/gami-common/dtos"
)

// UserType const
type UserType string

// AuthClaims represents jwt claims information.
type AuthClaims struct {
	jwt.StandardClaims
	UserID   string   `json:"id,omitempty"`
	UserType UserType `json:"id_type,omitempty"` // not use
}

var (
	authClaimsCtxKey  ctxKey = "gamification_id"
	authClaims0CtxKey ctxKey = "gamification_id_v2"
)

func SetAuth0Claim(parent context.Context, claims *dtos.Auth0Claims) context.Context {
	return setCtxValue(parent, authClaims0CtxKey, claims)
}

// SetAuthClaims set for backtrack version can be use without any break.
func SetAuthClaims(parent context.Context, claim0 *dtos.Auth0Claims) context.Context {
	return setCtxValue(parent, authClaimsCtxKey, &AuthClaims{
		StandardClaims: claim0.StandardClaims,
		UserID:         fmt.Sprintf("%v", claim0.VinIdInfo.UserID),
		UserType:       "",
	})
}

func GetAuth0Claims(ctx context.Context) *dtos.Auth0Claims {
	authClaims, ok := getCtxValue(ctx, authClaimsCtxKey).(*dtos.Auth0Claims)
	if ok {
		return authClaims
	}
	return nil
}

// GetAuthClaims gets jwt.AuthClaims in request's context.
func GetAuthClaims(ctx context.Context) *AuthClaims {
	authClaims, ok := getCtxValue(ctx, authClaimsCtxKey).(*AuthClaims)
	if ok {
		return authClaims
	}
	return &AuthClaims{}
}
