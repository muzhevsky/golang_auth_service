package v1

import (
	"authorization/controllers/http/middleware"
	"authorization/controllers/requests"
	_ "authorization/docs"
	"authorization/internal"
	"authorization/internal/errs"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInController struct {
	user   internal.ISignInUseCase
	logger logger.ILogger
}

func NewSignInController(handler *gin.Engine, useCase internal.ISignInUseCase, logger logger.ILogger) {
	u := &signInController{useCase, logger}

	handler.POST("/signin", u.signIn)
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
// @Router       /signin [post]
func (router *signInController) signIn(c *gin.Context) {
	var request requests.SignInRequest
	if err := c.ShouldBind(&request); err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	response, err := router.user.SignIn(c, &request)

	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}
