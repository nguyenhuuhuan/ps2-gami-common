package errors

import (
	"net/http"
)

// Module constants definition.
const (
	ModuleCommon = 00
)

// Common module error codes definition.
var (
	UnauthorizedCodeError = fmtErrorCode(http.StatusUnauthorized, ModuleCommon, 1)
	ErrNoResponse         = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 3)
	InternalServerError   = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 1)
)

func fmtErrorCode(status, module, code int) ErrorCode {
	return ErrorCode{
		status:     status,
		module:     module,
		detailCode: code,
	}
}
