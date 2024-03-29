package v1

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signUpRouter struct {
	user         usecase.IUser
	verification usecase.IVerification
	logger       logger.ILogger
}

func newSignUpRouter(handler *gin.RouterGroup, user usecase.IUser, verification usecase.IVerification, logger logger.ILogger) {
	u := &signUpRouter{user, verification, logger}

	handler.POST("/", u.signUp)
}

type createUserRequest struct {
	Login    string `json:"login" binding:"required" example:"TopPlayer123"`
	Password string `json:"password" binding:"required" example:"123superPassword"`
	EMail    string `json:"e-mail" binding:"required" example:"andrew123@qwerty.kom"`
	Nickname string `json:"nickname" binding:"required" example:"Looser1123"`
}
type createUserResponse struct {
	Id int `json:"id" example:"2"`
}

func (u *signUpRouter) signUp(c *gin.Context) {
	var userRequest createUserRequest

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
		} else if errors.Is(err, usecase.RecordAlreadyExists) {
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

	c.JSON(http.StatusOK, createUserResponse{user.Id})
	return
}
