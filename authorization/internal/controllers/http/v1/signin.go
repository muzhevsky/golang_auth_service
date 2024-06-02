package v1

import (
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInRouter struct {
	user   internal.ISignInUseCase
	logger logger.ILogger
}

func NewSignInRouter(handler *gin.Engine, useCase internal.ISignInUseCase, logger logger.ILogger) {
	u := &signInRouter{useCase, logger}

	handler.POST("/signin", u.signIn)
}

func (router *signInRouter) signIn(c *gin.Context) {
	var request requests.SignInRequest
	if err := c.ShouldBind(&request); err != nil {
		AddGinError(c, err)
		return
	}

	session, err := router.user.SignIn(c, &request)

	if err != nil {
		AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, session)
	return
}
