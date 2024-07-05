package v1

import (
	"authorization/controllers/http/middleware"
	"authorization/controllers/requests"
	_ "authorization/docs"
	"authorization/internal"
	"authorization/internal/errs"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type verificationController struct {
	verification internal.IVerifyUserUseCase
	logger       logger.ILogger
}

func NewVerificationController(handler *gin.Engine, verification internal.IVerifyUserUseCase, l logger.ILogger) {
	u := &verificationController{verification, l}

	handler.POST("/user/verify", u.verifyUser)
}

// UpdateById godoc
// @Summary      верификация пользователя
// @Description верификация пользователя с использованием токена, переданного в заголовке "Authorization"
// @Accept       json
// @Produce      json
// @Param request body requests.VerificationRequest true "request format"
// @Param Authorization header string true "access token"
// @Success      200  "Ok"
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "некорректный access token"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/verify [post]
func (u *verificationController) verifyUser(c *gin.Context) {
	var request requests.VerificationRequest
	if err := c.ShouldBind(&request); err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	userId, exists := c.Get("accountId")
	if !exists {
		err, _ := c.Get("authError")
		middleware.AddGinError(c, err.(error))
		return
	}

	err := u.verification.Verify(c, userId.(int), request.Code)

	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
