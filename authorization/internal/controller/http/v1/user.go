package v1

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRoutes struct {
	user         usecase.IUser
	verification usecase.IVerification
	l            logger.ILogger
}

func newUserRoutes(handler *gin.RouterGroup, user usecase.IUser, verification usecase.IVerification, l logger.ILogger) {
	u := &userRoutes{user, verification, l}

	handler.POST("/", u.createUser)
	handler.POST("/verify", u.verifyUser)
	handler.POST("/signin", u.signIn)
}

type createUserRequest struct {
	Login    string `json:"login" binding:"required" example:"TopPlayer123"`
	Password string `json:"password" binding:"required" example:"123superPassword"`
	EMail    string `json:"e-mail" binding:"required" example:"shilo@milo.psih"`
	Nickname string `json:"nickname" binding:"required" example:"Looser1123"`
}
type createUserResponse struct {
	Id int `json:"id" example:"2"`
}

// @Summary     Create user
// @ID          history
// @Accept      json
// @Produce     json
// @Param		request body createUserRequest true "Data for registration"
// @Success     200 {object} createUserResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Failure     409 {object} response
// @Router      /user/ [post]
func (u *userRoutes) createUser(c *gin.Context) {
	var userRequest createUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		u.l.Error(err, "http - v1 - createUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body", DefaultErrorCode)
		return
	}

	user, err := u.user.CreateUser(c, &entities.User{
		Login:    userRequest.Login,
		Password: userRequest.Password,
		EMail:    userRequest.EMail,
		Nickname: userRequest.Nickname,
	})

	if err != nil {
		u.l.Error(err, "http - v1 - createUser")
		if errors.Is(err, entities.ValidationError) {
			errorResponse(c, http.StatusBadRequest, err.Error(), ValidationErrorCode)
			return
		} else if errors.Is(err, usecase.RecordAlreadyExists) {
			errorResponse(c, http.StatusConflict, err.Error(), RecordExistErrorCode)
			return
		} else {
			errorResponse(c, http.StatusInternalServerError, err.Error(), DefaultErrorCode)
			return
		}
	}

	err = u.verification.CreateVerification(user)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error(), DefaultErrorCode)
	}

	c.JSON(http.StatusOK, createUserResponse{user.Id})
	return
}

type userVerificationRequest struct {
	UserId int    `json:"userId" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

func (u *userRoutes) verifyUser(c *gin.Context) {
	var request userVerificationRequest
	if err := c.ShouldBind(&request); err != nil {
		u.l.Error(err, "http - v1 - verifyUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body", DefaultErrorCode)
		return
	}

	success, err := u.verification.Verify(c,
		&entities.Verification{
			UserId: request.UserId,
			Code:   request.Code,
		})
	if err != nil {
		if errors.Is(err, entities.ExpiredCode) {
			errorResponse(c, http.StatusBadRequest, "code is outdated", DefaultErrorCode)
			return
		} else {
			u.l.Error(err, "http - v1 - verifyUser")
			errorResponse(c, http.StatusInternalServerError, "verification failed due to server couldn't handle the request", DefaultErrorCode)
			return
		}
	}

	if !success {
		errorResponse(c, http.StatusBadRequest, "invalid verification code", DefaultErrorCode)
		return
	}

	c.JSON(http.StatusOK, 321) // todo возвращать токены сессии
}

type userSignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *userRoutes) signIn(c *gin.Context) {
	var request userSignInRequest
	if err := c.ShouldBind(&request); err != nil {
		u.l.Error(err, "http - v1 - signIn")
		errorResponse(c, http.StatusBadRequest, "invalid request body", DefaultErrorCode)
		return
	}

	success, err := u.user.SignIn(c, &entities.User{
		Login:    request.Login,
		EMail:    request.Login, // todo подумать это вообще нормально?
		Password: request.Password,
	})

	if errors.Is(err, entities.UserIsNotVerified) {
		errorResponse(c, http.StatusUnauthorized, "user hasn't been verified yet", DefaultErrorCode)
		return
	}

	if err != nil { // TODO подумать над ошибками, если пользователь не существует
		errorResponse(c, http.StatusBadRequest, "user doesn't exist", DefaultErrorCode)
		return
	}

	if !success {
		errorResponse(c, http.StatusBadRequest, "password didn't match", DefaultErrorCode)
		return
	}

	c.JSON(http.StatusOK, 123) // todo возвращать токен сессии
}
