package v1

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInRouter struct {
	user usecase.IUser
	auth usecase.ISession
	l    logger.ILogger
}

type userSignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type userSignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func newSignInRouter(handler *gin.RouterGroup, user usecase.IUser, auth usecase.ISession, l logger.ILogger) {
	u := &signInRouter{user, auth, l}

	handler.POST("/signin", u.signIn)
}

func (router *signInRouter) signIn(c *gin.Context) {
	var request userSignInRequest
	if err := c.ShouldBind(&request); err != nil {
		router.l.Error(err, "http - v1 - signIn")
		errorResponse(c, http.StatusBadRequest, "invalid request body", DataBindErrorCode)
		return
	}

	user, err := router.user.SignIn(c, &entities.User{
		Login:    request.Login,
		EMail:    request.Login,
		Password: request.Password,
	})

	if err == nil {
		session, err := router.auth.CreateSession(c, user)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, "unexpected error", DefaultErrorCode)
			return
		}

		c.JSON(http.StatusOK, userSignInResponse{AccessToken: session.AccessToken, RefreshToken: session.RefreshToken})
		return
	}

	if errors.Is(err, entities.UserIsNotVerified) {
		errorResponse(c, http.StatusUnauthorized, "user hasn't been verified yet", UserIsNotVerifiedErrorCode)
		return
	}

	if errors.Is(err, entities.UserNotFound) {
		errorResponse(c, http.StatusBadRequest, "user doesn't exist", DefaultErrorCode)
		return
	}

	if errors.Is(err, entities.WrongPassword) {
		errorResponse(c, http.StatusBadRequest, "wrong password", DefaultErrorCode)
		return
	}

	errorResponse(c, http.StatusInternalServerError, "unexpected error", DefaultErrorCode)
}
