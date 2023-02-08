package errors_v2

import (
	"net/http"
)

const (
	ModuleCommon = 00
)

var (
	UnauthorizedCodeError = FmtErrorCode(http.StatusUnauthorized, ModuleCommon, 1, "Unauthorized")
	ErrNoResponse         = FmtErrorCode(http.StatusInternalServerError, ModuleCommon, 3, "No response error")
	InternalServerError   = FmtErrorCode(http.StatusInternalServerError, ModuleCommon, 1, "Internal server error")
)
