package grpc

import (
	"authorization/internal/controllers/requests"
	"authorization/internal/entities"
	"authorization/internal/usecases"
	"authorization/pkg/grpc/proto"
	"authorization/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerController struct {
	user         usecases.IUser
	verification usecases.IVerification
	logger       logger.ILogger
}

func newRegisterController(server *proto.AuthServer, user usecases.IUser, verification usecases.IVerification, logger logger.ILogger) *registerController {
	return &registerController{user, verification, logger}
}

func (u *registerController) register(c *gin.Context) {
	var userRequest requests.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		u.logger.Error(err, "http - v1 - createUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body", DataBindErrorCode)
		return
	}

	user, err := u.user.CreateUser(c, &entities.User{
		Login:    userRequest.Login,
		Password: userRequest.Password,
		EMail:    userRequest.EMail,
		Nickname: userRequest.Nickname,
	})

	if err != nil {
		if errors.Is(err, entities.ValidationError) {
			errorResponse(c, http.StatusBadRequest, err.Error(), LoginValidationErrorCode)
			return
		} else if errors.Is(err, usecases.RecordAlreadyExists) {
			errorResponse(c, http.StatusConflict, err.Error(), RecordExistErrorCode)
			return
		} else {
			u.logger.Error(err, "http - v1 - createUser")
			errorResponse(c, http.StatusInternalServerError, "Internal server error", DefaultErrorCode)
			return
		}
	}

	err = u.verification.CreateVerification(c, user)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error(), DefaultErrorCode)
	}

	c.JSON(http.StatusOK, requests.CreateUserResponse{user.Id})
	return
}
