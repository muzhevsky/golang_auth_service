package v1

import "github.com/gin-gonic/gin"

const DefaultErrorCode = "D-0000"

const DataBindErrorCode = "Client-0001"

const LoginValidationErrorCode = "Validation-0001"
const EmailValidationErrorCode = "Validation-0002"
const PasswordValidationErrorCode = "Validation-0003"

const VerificationExpiredErrorCode = "Verification-0001"
const WrongVerificationErrorCode = "Verification-0002"
const UserIsNotVerifiedErrorCode = "Verification-0003"

const RecordExistErrorCode = "R-0001"

type response struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func errorResponse(c *gin.Context, code int, msg string, developerCode string) {
	c.AbortWithStatusJSON(code, response{msg, developerCode})
}
