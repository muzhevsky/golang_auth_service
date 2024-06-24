package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"smartri_app/internal/errs"
	"smartri_app/pkg/logger"
)

type ErrorHandler struct {
	logger logger.ILogger
}

func NewErrorHandler(logger logger.ILogger) *ErrorHandler {
	return &ErrorHandler{logger: logger}
}

func (h *ErrorHandler) HandleError(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		err := c.Errors.Last()

		h.logger.Debug(err)

		// Common ////////////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.DataBindError) {
			response(c, http.StatusBadRequest, err.Error(), DataBindErrorCode)
			return
		}
		if errors.Is(err, errs.UnauthenticatedError) {
			response(c, http.StatusUnauthorized, err.Error(), UnauthenticatedErrorCode)
			return
		}
		///////////////////////////////////////////////////////////////////////////////////

		response(c, http.StatusInternalServerError, "Internal server error", InternalServerErrorErrorCode)
	}
}

type ErrorResponse struct {
	Error string `json:"error" example:"Some error description"`
	Code  string `json:"code" example:"ErrorGroup-0001"`
}

func response(c *gin.Context, code int, msg string, developerCode string) {
	c.AbortWithStatusJSON(code, ErrorResponse{msg, developerCode})
}
