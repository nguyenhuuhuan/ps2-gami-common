package errors_v2

import (
	"fmt"
	"strconv"
)

type ErrorCode struct {
	status     int
	module     int
	detailCode int
	message    string
}

func (errCode ErrorCode) Code() int {
	errStr := fmt.Sprintf("%d%02d%02d", errCode.status, errCode.module, errCode.detailCode)
	code, _ := strconv.Atoi(errStr)
	return code
}

func (errCode ErrorCode) Status() int {
	return errCode.status
}

func (errCode ErrorCode) Module() int {
	return errCode.module
}

func (errCode ErrorCode) DetailCode() int {
	return errCode.detailCode
}

func (errCode ErrorCode) Message() string {
	return errCode.message
}

type AppError struct {
	Meta          ErrorMeta `json:"meta"`
	OriginalError error     `json:"-"`
	ErrorCode     ErrorCode `json:"-"`
}

type ErrorMeta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (appErr AppError) Error() string {
	if appErr.OriginalError != nil {
		return appErr.OriginalError.Error()
	}
	return appErr.Meta.Message
}

func New(errCode ErrorCode) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: errCode.Message(),
		},
		OriginalError: nil,
		ErrorCode:     errCode,
	}
}

func NewWithMsg(errCode ErrorCode, msg string) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: msg,
		},
		OriginalError: nil,
		ErrorCode:     errCode,
	}
}

func FmtErrorCode(status, module, code int, message string) ErrorCode {
	return ErrorCode{
		status:     status,
		module:     module,
		detailCode: code,
		message:    message,
	}
}

func NewErrorCode(code int, message string) ErrorCode {
	return ErrorCode{
		status:     code / 10000,
		module:     (code / 100) % 100,
		detailCode: code % 100,
		message:    message,
	}
}
