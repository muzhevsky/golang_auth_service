package v1

import (
	"authorization/controllers/http/middleware"
	"authorization/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAccountDevices struct {
	useCase internal.IGetAccountDevicesUseCase
}

func NewGetAccountDevices(handler *gin.Engine, useCase internal.IGetAccountDevicesUseCase) {
	g := &getAccountDevices{useCase: useCase}

	handler.GET("/user/devices", g.GetAccountDevices)
}

// GetAccountDevices godoc
// @Summary      запрос на получение устройств пользователя
// @Description  запрос на получение устройств, на которых существует открытая сессия данного пользователя пользователя
// с использованием токена, переданного в заголовке "Authorization"
// @Produce      json
// @Param Authorization header string true "access token"
// @Success 200 {object} requests.AccountDevicesResponse
// @Failure 401 {object} middleware.ErrorResponse "некорректный access token"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/devices [get]
func (con *getAccountDevices) GetAccountDevices(c *gin.Context) {
	accountId, exists := c.Get("accountId")
	if !exists {
		err, _ := c.Get("authError")
		middleware.AddGinError(c, err.(error))
		return
	}

	id := accountId.(int)

	response, err := con.useCase.GetAccountDevices(c, id)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
