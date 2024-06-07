package v1

import (
	_ "authorization/docs"
	"authorization/internal"
	http2 "authorization/internal/controllers/http/middleware"
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
// @Summary      вход в аккаунт
// @Description  вход в аккаунт с использованием пар логин + пароль или email + пароль для получения токенов
// @Accept       json
// @Produce      json
// @Param request body requests.SignInRequest true "структура запроса"
// @Success      200  {object}  requests.SignInResponse
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "неправильный пароль"
// @Failure 404 {object} middleware.ErrorResponse "пользователь не найден"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /auth/signin [post]
func (router *signInRouter) signIn(c *gin.Context) {
	var request requests.SignInRequest
	if err := c.ShouldBind(&request); err != nil {
		http2.AddGinError(c, errs.DataBindError)
		return
	}

	session, err := router.user.SignIn(c, &request)

	if err != nil {
		http2.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.SignInResponse{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt.Unix(),
	})
	return
}
