package v1

import (
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type sessionRoutes struct {
	session usecase.ISession
	l       logger.ILogger
}

func newSessionRoutes(handler *gin.RouterGroup, session usecase.ISession, l logger.ILogger) {
	u := &sessionRoutes{session, l}

	handler.GET("/", u.authenticateUser)
}

type authenticateUserRequest struct {
	Login    string
	Password string
}
type authenticateUserResponse struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     time.Time
}

func (u *sessionRoutes) authenticateUser(c *gin.Context) {
	var authRequest authenticateUserRequest

	if err := c.ShouldBindJSON(&authRequest); err != nil {
		u.l.Error(err, "http - v1 - authenticateUser")
		errorResponse(c, http.StatusBadRequest, "invalid request body", DefaultErrorCode)
		return
	}

	_, err := u.session.AuthenticateUser(authRequest.Login, authRequest.Password)
	if err != nil {
		//u.l.Error(err, "http - v1 - authenticateUser")
		//if errors.Is(err, entities.ValidationError) {
		//	errorResponse(c, http.StatusBadRequest, err.Error(), ValidationErrorCode)
		//	return
		//} else if errors.Is(err, usecase.RecordAlreadyExists) {
		//	errorResponse(c, http.StatusConflict, err.Error(), RecordExistErrorCode)
		//	return
		//} else {
		//	errorResponse(c, http.StatusInternalServerError, err.Error(), DefaultErrorCode)
		//	return
		//}
	}
	c.JSON(http.StatusOK, authenticateUserResponse{})
	return
}
