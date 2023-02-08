package ctxutil

import (
	"context"
	"strconv"
)

type ctxKey string

var userIDCtxKey ctxKey = "user_id"
var errorCodeCtxKey ctxKey = "error_code"

// SetUserID sets dtos.UserID to context of a request.
func SetUserID(parent context.Context, userID int64) context.Context {
	return setCtxValue(parent, userIDCtxKey, userID)
}

// GetUserID gets dtos.UserID in request's context.
func GetUserID(ctx context.Context) int64 {
	authClaims, ok := getCtxValue(ctx, userIDCtxKey).(int64)
	if ok {
		return authClaims
	}
	return 0
}

func SetUserIDStr(parent context.Context, userID string) context.Context {
	return setCtxValue(parent, userIDCtxKey, userID)
}

// GetUserIDStr gets dtos.UserID in request's context.
func GetUserIDStr(ctx context.Context) string {
	authClaims, ok := getCtxValue(ctx, userIDCtxKey).(string)
	if ok {
		return authClaims
	}

	authNumber, ok := getCtxValue(ctx, userIDCtxKey).(int64)
	if ok {
		return strconv.FormatInt(authNumber, 10)
	}
	return ""
}

func setCtxValue(parent context.Context, key ctxKey, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}

func getCtxValue(ctx context.Context, key ctxKey) interface{} {
	if ctx != nil {
		return ctx.Value(key)
	}
	return nil
}

func SetErrorCode(parent context.Context, errorCode int) context.Context {
	return setCtxValue(parent, errorCodeCtxKey, errorCode)
}

func GetErrorCode(ctx context.Context) int {
	errorCode, ok := getCtxValue(ctx, errorCodeCtxKey).(int)
	if ok {
		return errorCode
	}
	return 0
}

func SetCtxValueString(parent context.Context, key ctxKey, value string) context.Context {
	return setCtxValue(parent, key, value)
}

func GetCtxValueString(ctx context.Context, key string) string {
	newContextKey := ctxKey(key)
	value, ok := getCtxValue(ctx, newContextKey).(string)
	if ok {
		return value
	}
	return ""
}
