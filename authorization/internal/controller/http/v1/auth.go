package v1

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type authRouter struct {
	user usecase.IUser
	auth usecase.ISession
	l    logger.ILogger
}

func newAuthRouter(handler *gin.RouterGroup, user usecase.IUser, auth usecase.ISession, l logger.ILogger) {
	u := &authRouter{user, auth, l}

	handler.POST("/auth", u.authorize)
}

type authRequest struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (r *authRouter) authorize(c *gin.Context) {
	request := authRequest{}
	if err := c.ShouldBind(&request); err != nil {
		r.l.Error(err, "http - v1 - verifyUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body", DataBindErrorCode)
		return
	}
	response := authRequest{}
	_, err := r.auth.VerifyAccessToken(c, request.AccessToken)
	if err != nil {
		if errors.Is(err, entities.AccessTokenExpired) {
			updatedSession, updateErr := r.auth.UpdateSession(c, &entities.Session{0, "", request.AccessToken, request.RefreshToken, time.Now()})
			if updateErr == nil {
				response.AccessToken = updatedSession.AccessToken
				response.RefreshToken = updatedSession.RefreshToken
			}
			err = updateErr
		}
		if err != nil {
			if errors.Is(err, entities.RefreshTokenExpired) ||
				errors.Is(err, entities.NotAValidAccessToken) ||
				errors.Is(err, entities.NotAValidRefreshToken) {
				errorResponse(c, http.StatusUnauthorized, err.Error(), DefaultErrorCode)
				return
			}

			r.l.Error("", err.Error())
			errorResponse(c, http.StatusUnauthorized, "unexpected error", DefaultErrorCode)
			return
		}
	}
	c.JSON(http.StatusOK, response)
}
