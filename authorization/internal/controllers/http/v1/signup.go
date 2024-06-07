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

type registerRouter struct {
	user         internal.ICreateUserUseCase
	verification internal.IVerifyUserUseCase
	logger       logger.ILogger
}

func NewSignUpRouter(handler *gin.Engine, user internal.ICreateUserUseCase, verification internal.IVerifyUserUseCase, logger logger.ILogger) {
	u := &registerRouter{user, verification, logger}

	handler.POST("/auth/signup", u.register)
}

// SignUp godoc
// @Summary      регистрация нового пользователя
// @Description  регистрация нового пользователя
// @Accept       json
// @Produce      json
// @Param request body requests.CreateUserRequest true "структура запрос"
// @Success      200  {object}  requests.CreateUserResponse
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 409 {object} middleware.ErrorResponse "пользователь уже существует"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /auth/signup [post]
func (u *registerRouter) register(c *gin.Context) {
	var userRequest requests.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		http2.AddGinError(c, errs.DataBindError)
		return
	}

	response, err := u.user.CreateUser(c, &userRequest)

	if err != nil {
		http2.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
