package v1

import (
	"authorization/controllers/http/middleware"
	_ "authorization/docs"
	"authorization/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type checkVerificationController struct {
	useCase internal.ICheckVerificationUseCase
}

func NewCheckVerificationController(handler *gin.Engine, useCase internal.ICheckVerificationUseCase) {
	u := &checkVerificationController{useCase: useCase}

	handler.GET("/user/verification", u.checkVerification)
}

// CheckVerification godoc
// @Summary      запрос на проверку верификации пользователя
// @Description  запрос на проверку верификации пользователя с использованием токена, переданного в заголовке "Authorization"
// @Produce      json
// @Param Authorization header string true "access token"
// @Success 200 {object} requests.CheckVerificationResponse
// @Failure 401 {object} middleware.ErrorResponse "некорректный access token"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/verification [get]
func (v *checkVerificationController) checkVerification(c *gin.Context) {
	accountId, exists := c.Get("accountId")
	if !exists {
		err, _ := c.Get("authError")
		middleware.AddGinError(c, err.(error))
		return
	}

	checked, err := v.useCase.CheckVerification(c, accountId.(int))
	if err != nil {
		middleware.AddGinError(c, err.(error))
		return
	}

	c.JSON(http.StatusOK, checked)
}
