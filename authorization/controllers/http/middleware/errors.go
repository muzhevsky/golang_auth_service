package middleware

import "github.com/gin-gonic/gin"

// TODO create approptiate errors that make sense
const InternalServerErrorErrorCode = "Unexpected"

const DataBindErrorCode = "Auth-Client-001"
const RecordExistErrorCode = "Auth-Client-002"
const LoginValidationErrorCode = "Auth-Client-003"

const AuthAccountNotFoundErrorCode = "Auth-Auth-001"
const AuthWrongPasswordErrorCode = "Auth-Auth-002"
const AuthSessionErrorCode = "Auth-Auth-003"
const RefreshSessionErrorCode = "Auth-Auth-004"

const VerificationExpiredErrorCode = "Auth-Verification-001"
const WrongVerificationErrorCode = "Auth-Verification-002"
const UserIsNotVerifiedErrorCode = "Auth-Verification-003"
const UserIsAlreadyVerifiedErrorCode = "Auth-Verification-004"

func AddGinError(ctx *gin.Context, err error) {
	ctx.Errors = append(ctx.Errors, &gin.Error{
		Err:  err,
		Type: gin.ErrorTypePublic,
		Meta: nil,
	})
}
