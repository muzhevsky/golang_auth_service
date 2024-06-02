package v1

import (
	"authorization/internal"
	"authorization/internal/controllers/requests"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type verificationRoute struct {
	verification internal.IVerification
	logger       logger.ILogger
}

func NewVerificationRouter(handler *gin.Engine, verification internal.IVerification, l logger.ILogger) {
	u := &verificationRoute{verification, l}

	handler.POST("/verify", u.verifyUser)
}

func (u *verificationRoute) verifyUser(c *gin.Context) {
	var request requests.VerificationRequest
	if err := c.ShouldBind(&request); err != nil {
		AddGinError(c, err)
		return
	}

	err := u.verification.Verify(c, &request)

	if err != nil {
		AddGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, "")
}
