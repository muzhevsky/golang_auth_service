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

type signupController struct {
	user         internal.ICreateAccountUseCase
	verification internal.IVerifyUserUseCase
	logger       logger.ILogger
}

func NewSignUpController(handler *gin.Engine, user internal.ICreateAccountUseCase, verification internal.IVerifyUserUseCase, logger logger.ILogger) {
	u := &signupController{user, verification, logger}

	handler.POST("/signup", u.signup)
}

// SignUp godoc
// @Summary      регистрация нового пользователя
// @Description  регистрация нового пользователя
// @Accept       json
// @Produce      json
// @Param request body requests.SignUpRequest true "структура запрос"
// @Success      200  {object}  requests.SignUpResponse
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 409 {object} middleware.ErrorResponse "пользователь уже существует"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /signup [post]
func (u *signupController) signup(c *gin.Context) {
	var userRequest requests.SignUpRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	response, err := u.user.CreateAccount(c, &userRequest)

	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
