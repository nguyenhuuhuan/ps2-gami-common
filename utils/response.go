package utils

import (
	"encoding/json"
	"net/http"

	"gitlab.id.vin/gami/gami-common/utils/ctxutil"

	"gitlab.id.vin/gami/gami-common/errors"

	"github.com/gin-gonic/gin"
)

// JSON responds a HTTP request with JSON data.
func JSON(c *gin.Context, data interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(errors.ErrNoResponse.Status(), errors.New(errors.ErrNoResponse))
	}
}

// HandleError handles error of HTTP request.
func HandleError(c *gin.Context, err error) {
	if err != nil {
		appErr, ok := err.(errors.AppError)
		if ok {
			newContext := ctxutil.SetErrorCode(c.Request.Context(), appErr.Meta.Code)
			c.Request = c.Request.WithContext(newContext)
			c.JSON(appErr.ErrorCode.Status(), appErr)
		} else {
			c.JSON(errors.InternalServerError.Status(), errors.New(errors.InternalServerError))
		}
	} else {
		c.JSON(errors.ErrNoResponse.Status(), errors.New(errors.ErrNoResponse))
	}
}

// HandleErrorWithoutContext return error response without context
func HandleErrorWithoutContext(err error) string {
	appErr, ok := err.(errors.AppError)
	if ok {
		data, _ := json.Marshal(appErr)
		return string(data)
	}

	return "internal server error"
}
