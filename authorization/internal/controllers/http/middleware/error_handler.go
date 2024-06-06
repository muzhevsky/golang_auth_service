package middleware

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

		// Common ////////////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.DataBindError) {
			response(c, http.StatusBadRequest, err.Error(), v1.DataBindErrorCode)
			return
		}

		if errors.Is(err, errs.RecordAlreadyExists) {
			response(c, http.StatusConflict, err.Error(), v1.RecordExistErrorCode)
			return
		}
		//////////////////////////////////////////////////////////////////////////////////

		// Validation ////////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.LoginValidationError) {
			response(c, http.StatusBadRequest, err.Error(), v1.LoginValidationErrorCode)
			return
		}
		if errors.Is(err, errs.EmailValidationError) {
			response(c, http.StatusBadRequest, err.Error(), v1.EmailValidationErrorCode)
			return
		}
		if errors.Is(err, errs.PasswordValidationError) {
			response(c, http.StatusBadRequest, err.Error(), v1.PasswordValidationErrorCode)
			return
		}
		///////////////////////////////////////////////////////////////////////////////////

		// Auth ///////////////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.WrongPassword) {
			response(c, http.StatusUnauthorized, err.Error(), v1.AuthWrongPasswordErrorCode)
			return
		}
		if errors.Is(err, errs.UserNotFound) {
			response(c, http.StatusNotFound, err.Error(), v1.AuthUserNotFoundErrorCode)
			return
		}
		///////////////////////////////////////////////////////////////////////////////////

		// Verification ///////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.WrongVerificationCode) {
			response(c, http.StatusBadRequest, err.Error(), v1.WrongVerificationErrorCode)
			return
		}
		if errors.Is(err, errs.NotAValidAccessToken) {
			response(c, http.StatusUnauthorized, err.Error(), v1.AuthSessionErrorCode)
			return
		}
		if errors.Is(err, errs.ExpiredVerificationCode) {
			response(c, http.StatusUnauthorized, err.Error(), v1.VerificationExpiredErrorCode)
			return
		}
		if errors.Is(err, errs.UserIsNotVerified) {
			response(c, http.StatusForbidden, err.Error(), v1.UserIsNotVerifiedErrorCode)
			return
		}
		if errors.Is(err, errs.UserIsAlreadyVerified) {
			response(c, http.StatusConflict, err.Error(), v1.UserIsAlreadyVerifiedErrorCode)
			return
		}
		///////////////////////////////////////////////////////////////////////////////////

		response(c, http.StatusInternalServerError, "Internal server error", v1.InternalServerErrorErrorCode)
	}
}

type ErrorResponse struct {
	Error string `json:"error" example:"Some error description"`
	Code  string `json:"code" example:"ErrorGroup-0001"`
}

func response(c *gin.Context, code int, msg string, developerCode string) {
	c.AbortWithStatusJSON(code, ErrorResponse{msg, developerCode})
}
