package middleware

import "github.com/gin-gonic/gin"

const InternalServerErrorErrorCode = "Unexpected"

const DataBindErrorCode = "Client-0001"
const UnauthenticatedErrorCode = "Auth-0005"

func AddGinError(ctx *gin.Context, err error) {
	ctx.Errors = append(ctx.Errors, &gin.Error{
		Err:  err,
		Type: gin.ErrorTypePublic,
		Meta: nil,
	})
}
