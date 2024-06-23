package v1

import (
	_ "authorization/docs"
	"authorization/internal"
	http2 "authorization/internal/controllers/http/middleware"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type requestVerificationController struct {
	user         internal.ICreateAccountUseCase
	verification internal.IRequestVerificationUseCase
	logger       logger.ILogger
}

func NewRequestVerificationRouter(handler *gin.Engine, user internal.ICreateAccountUseCase, verification internal.IRequestVerificationUseCase, logger logger.ILogger) {
	u := &requestVerificationController{user, verification, logger}

	handler.POST("/verification/request", u.requestVerification)
}

// RequestVerification godoc
// @Summary      запрос на верификацию пользователя
// @Description запрос на верификацию пользователя с использованием токена, переданного в заголовке "Authorization"
// @Accept       json
// @Produce      json
// @Success      200  "(TODO: в данные момент возвращается код, чтобы каждый раз клиент не мучал почту) Ok"
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "некорректный access token"
// @Failure 409 {object} middleware.ErrorResponse "пользователь уже верифицирован"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /verification/request [post]
func (u *requestVerificationController) requestVerification(c *gin.Context) {
	userId, exists := c.Get("accountId")
	if !exists {
		err, _ := c.Get("authError")
		http2.AddGinError(c, err.(error))
		return
	}

	code, err := u.verification.RequestVerification(c, userId.(int))
	if err != nil {
		http2.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, code)
}
