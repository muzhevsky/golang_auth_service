package v1

import (
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerRouter struct {
	user         internal.ICreateUserUseCase
	verification internal.IVerification
	logger       logger.ILogger
}

func NewSignUpRouter(handler *gin.Engine, user internal.ICreateUserUseCase, verification internal.IVerification, logger logger.ILogger) {
	u := &registerRouter{user, verification, logger}

	handler.POST("/signup", u.register)
}

func (u *registerRouter) register(c *gin.Context) {
	var userRequest requests.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		AddGinError(c, err)
		return
	}

	user, err := u.user.CreateUser(c, &userRequest)

	if err != nil {
		AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, requests.CreateUserResponse{user.Id})
}
