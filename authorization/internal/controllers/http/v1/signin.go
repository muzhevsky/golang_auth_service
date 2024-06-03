package v1

import (
	_ "authorization/docs"
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/internal/errs"
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

	handler.POST("/auth/signin", u.signIn)
}

// SignIn godoc
// @Summary      sign in
// @Description  sign in
// @Accept       json
// @Produce      json
// @Param request body requests.SignInRequest true "request format"
// @Success      200  {object}  requests.SignInResponse
// @Router       /auth/signin [post]
func (router *signInRouter) signIn(c *gin.Context) {
	var request requests.SignInRequest
	if err := c.ShouldBind(&request); err != nil {
		AddGinError(c, errs.DataBindError)
		return
	}

	session, err := router.user.SignIn(c, &request)

	if err != nil {
		AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.SignInResponse{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		ExpireAt:     session.ExpireAt.Unix(),
	})
	return
}
