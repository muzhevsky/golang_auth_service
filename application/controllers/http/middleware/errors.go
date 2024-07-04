package middleware

import "github.com/gin-gonic/gin"

const InternalServerErrorErrorCode = "Unexpected"

const DataBindErrorCode = "Core-Client-001"

const UserDataNotFoundErrorCode = "Core-User-001"
const UserAvatarNotFoundErrorCode = "Core-User-002"

const TestIsAlreadyPassedErrorCode = "Core-Test-001"

const UnauthenticatedErrorCode = "Core-Auth-001"

func AddGinError(ctx *gin.Context, err error) {
	ctx.Errors = append(ctx.Errors, &gin.Error{
		Err:  err,
		Type: gin.ErrorTypePublic,
		Meta: nil,
	})
}
