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
// @Summary      creates user
// @Description  creates user
// @Accept       json
// @Produce      json
// @Param request body requests.CreateUserRequest true "request format"
// @Success      200  {object}  requests.CreateUserResponse
// @Router       /auth/signup [post]
func (u *registerRouter) register(c *gin.Context) {
	var userRequest requests.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		AddGinError(c, errs.DataBindError)
		return
	}

	user, err := u.user.CreateUser(c, &userRequest)

	if err != nil {
		AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.CreateUserResponse{user.Id})
}
