package http

import (
	v1 "authorization/internal/controllers/http/v1"
	errs "authorization/internal/errs"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		err := c.Errors.Last()

		fmt.Println(err) // todo поменять логгер

		if errors.Is(err, errs.ValidationError) {
			response(c, http.StatusBadRequest, err.Error(), v1.LoginValidationErrorCode)
			return
		}
		if errors.Is(err, errs.DataBindError) {
			response(c, http.StatusBadRequest, err.Error(), v1.DataBindErrorCode)
			return
		}
		if errors.Is(err, errs.WrongVerificationCode) {
			response(c, http.StatusBadRequest, err.Error(), v1.DataBindErrorCode)
			return
		}
		if errors.Is(err, errs.RecordAlreadyExists) {
			response(c, http.StatusConflict, err.Error(), v1.RecordExistErrorCode)
			return
		}
		if errors.Is(err, errs.UserNotFound) {
			response(c, http.StatusBadRequest, err.Error(), v1.InternalServerErrorErrorCode) // todo сделать код ошибки
			return
		}

		response(c, http.StatusInternalServerError, "Internal server error", v1.InternalServerErrorErrorCode)
	}
}

type errorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func response(c *gin.Context, code int, msg string, developerCode string) {
	c.AbortWithStatusJSON(code, errorResponse{msg, developerCode})
}
