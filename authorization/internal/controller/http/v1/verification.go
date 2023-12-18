package v1

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type verificationRoute struct {
	verification usecase.IVerification
	auth         usecase.ISession
	l            logger.ILogger
}

func newVerificationRoute(handler *gin.RouterGroup, verification usecase.IVerification, auth usecase.ISession, l logger.ILogger) {
	u := &verificationRoute{verification, auth, l}

	handler.POST("/verify", u.verifyUser)
}

type userVerificationRequest struct {
	UserId int    `json:"userId" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

func (u *verificationRoute) verifyUser(c *gin.Context) {
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

	c.JSON(http.StatusOK, "")
}
