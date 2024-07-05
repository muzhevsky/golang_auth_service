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

type refreshSessionController struct {
	useCase internal.IRefreshSessionUseCase
	logger  logger.ILogger
}

func NewRefreshSessionController(handler *gin.Engine, useCase internal.IRefreshSessionUseCase, logger logger.ILogger) {
	u := &refreshSessionController{
		useCase: useCase,
		logger:  logger,
	}

	handler.POST("/token/update", u.refreshSession)
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
// @Router       /token/update [post]
func (r *refreshSessionController) refreshSession(c *gin.Context) {
	request := requests.RefreshSessionRequest{}
	if err := c.ShouldBind(&request); err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}
	response := requests.RefreshSessionResponse{}
	session, err := r.useCase.RefreshSession(c, &request)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	response.AccessToken = session.AccessToken
	response.RefreshToken = session.RefreshToken
	c.JSON(http.StatusOK, response)
}
