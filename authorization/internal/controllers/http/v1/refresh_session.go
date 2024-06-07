package v1

import (
	_ "authorization/docs"
	"authorization/internal"
	http2 "authorization/internal/controllers/http/middleware"
	"authorization/internal/controllers/requests"
	"authorization/internal/errs"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type refreshSessionRouter struct {
	useCase internal.IRefreshSessionUseCase
	logger  logger.ILogger
}

func NewRefreshSessionRouter(handler *gin.Engine, useCase internal.IRefreshSessionUseCase, logger logger.ILogger) {
	u := &refreshSessionRouter{
		useCase: useCase,
		logger:  logger,
	}

	handler.POST("/auth/token/update", u.refreshSession)
}

// RefreshSession godoc
// @Summary      обновление сессии
// @Description  возвращает новую пару токенов при отправке старой пары и при условии их валидности
// @Accept       json
// @Produce      json
// @Param request body requests.RefreshSessionRequest true "request format"
// @Success      200  {object}  requests.RefreshSessionResponse
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "невалидная пара токенов, либо истекший refresh token"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /auth/token/update [post]
func (r *refreshSessionRouter) refreshSession(c *gin.Context) {
	request := requests.RefreshSessionRequest{}
	if err := c.ShouldBind(&request); err != nil {
		http2.AddGinError(c, errs.DataBindError)
		return
	}
	response := requests.RefreshSessionResponse{}
	session, err := r.useCase.RefreshSession(c, &request)
	if err != nil {
		http2.AddGinError(c, err)
		return
	}

	response.AccessToken = session.AccessToken
	response.RefreshToken = session.RefreshToken
	c.JSON(http.StatusOK, response)
}
