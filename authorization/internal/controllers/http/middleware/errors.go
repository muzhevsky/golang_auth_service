package middleware

import "github.com/gin-gonic/gin"

const InternalServerErrorErrorCode = "Unexpected"

const DataBindErrorCode = "Client-0001"
const RecordExistErrorCode = "Client-0002"

const AuthUserNotFoundErrorCode = "Auth-0001"
const AuthWrongPasswordErrorCode = "Auth-0002"
const AuthSessionErrorCode = "Auth-0003"
const RefreshSessionErrorCode = "Auth-0004"

const LoginValidationErrorCode = "Validation-0001"
const EmailValidationErrorCode = "Validation-0002"
const PasswordValidationErrorCode = "Validation-0003"

const VerificationExpiredErrorCode = "Verification-0001"
const WrongVerificationErrorCode = "Verification-0002"
const UserIsNotVerifiedErrorCode = "Verification-0003"
const UserIsAlreadyVerifiedErrorCode = "Verification-0004"

func AddGinError(ctx *gin.Context, err error) {
	ctx.Errors = append(ctx.Errors, &gin.Error{
		Err:  err,
		Type: gin.ErrorTypePublic,
		Meta: nil,
	})
}
