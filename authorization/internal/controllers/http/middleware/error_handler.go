package middleware

import (
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
			response(c, http.StatusBadRequest, err.Error(), DataBindErrorCode)
			return
		}

		if errors.Is(err, errs.RecordAlreadyExists) {
			response(c, http.StatusConflict, err.Error(), RecordExistErrorCode)
			return
		}
		//////////////////////////////////////////////////////////////////////////////////

		// Validation ////////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.LoginValidationError) {
			response(c, http.StatusBadRequest, err.Error(), LoginValidationErrorCode)
			return
		}
		if errors.Is(err, errs.EmailValidationError) {
			response(c, http.StatusBadRequest, err.Error(), EmailValidationErrorCode)
			return
		}
		if errors.Is(err, errs.PasswordValidationError) {
			response(c, http.StatusBadRequest, err.Error(), PasswordValidationErrorCode)
			return
		}
		///////////////////////////////////////////////////////////////////////////////////

		// Auth ///////////////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.WrongPassword) {
			response(c, http.StatusUnauthorized, err.Error(), AuthWrongPasswordErrorCode)
			return
		}
		if errors.Is(err, errs.UserNotFound) {
			response(c, http.StatusNotFound, err.Error(), AuthUserNotFoundErrorCode)
			return
		}
		if errors.Is(err, errs.NotAValidAccessToken) || errors.Is(err, errs.AccessTokenExpired) {
			response(c, http.StatusUnauthorized, err.Error(), AuthSessionErrorCode)
			return
		}
		if errors.Is(err, errs.NotAValidRefreshToken) || errors.Is(err, errs.RefreshTokenExpired) {
			response(c, http.StatusUnauthorized, err.Error(), RefreshSessionErrorCode)
			return
		}
		///////////////////////////////////////////////////////////////////////////////////

		// Verification ///////////////////////////////////////////////////////////////////
		if errors.Is(err, errs.WrongVerificationCode) {
			response(c, http.StatusBadRequest, err.Error(), WrongVerificationErrorCode)
			return
		}
		if errors.Is(err, errs.ExpiredVerificationCode) {
			response(c, http.StatusUnauthorized, err.Error(), VerificationExpiredErrorCode)
			return
		}
		if errors.Is(err, errs.UserIsNotVerified) {
			response(c, http.StatusForbidden, err.Error(), UserIsNotVerifiedErrorCode)
			return
		}
		if errors.Is(err, errs.UserIsAlreadyVerified) {
			response(c, http.StatusConflict, err.Error(), UserIsAlreadyVerifiedErrorCode)
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
