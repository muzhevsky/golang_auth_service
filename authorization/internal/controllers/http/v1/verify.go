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

type verificationRoute struct {
	verification internal.IVerifyUserUseCase
	logger       logger.ILogger
}

func NewVerificationRouter(handler *gin.Engine, verification internal.IVerifyUserUseCase, l logger.ILogger) {
	u := &verificationRoute{verification, l}

	handler.POST("/user/verify", u.verifyUser)
}

// Verify godoc
// @Summary      verifies user
// @Description  AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAaa
// @Accept       json
// @Produce      json
// @Param request body requests.VerificationRequest true "request format"
// @Success      200  "Ok"
// @Router       /user/verify [post]
func (u *verificationRoute) verifyUser(c *gin.Context) {
	var request requests.VerificationRequest
	if err := c.ShouldBind(&request); err != nil {
		AddGinError(c, errs.DataBindError)
		return
	}

	err := u.verification.Verify(c, &request)

	if err != nil {
		AddGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
