package v1

import (
	"authorization/controllers/http/middleware"
	"authorization/controllers/requests"
	"authorization/internal"
	"authorization/internal/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type closeSessionsController struct {
	useCase internal.ICloseSessionsByIdsUseCase
}

func NewCloseSessionsController(handler *gin.Engine, useCase internal.ICloseSessionsByIdsUseCase) {
	g := &closeSessionsController{useCase: useCase}

	handler.POST("/user/sessions/close", g.CloseSessions)
}

// CloseSessions godoc
// @Summary      запрос на закрытие сессий пользователя
// @Description  запрос на закрытие сессий пользователя по их id с использованием токена, переданного в заголовке "Authorization"
// @Produce      json
// @Param Authorization header string true "access token"
// @Param request body requests.CloseSessionsRequest true "request format"
// @Success 200 "ok"
// @Failure 401 {object} middleware.ErrorResponse "некорректный access token"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/sessions/close [post]
func (con *closeSessionsController) CloseSessions(c *gin.Context) {
	accountId, exists := c.Get("accountId")
	if !exists {
		err, _ := c.Get("authError")
		middleware.AddGinError(c, err.(error))
		return
	}

	id := accountId.(int)

	request := requests.CloseSessionsRequest{}
	if err := c.ShouldBind(&request); err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	err := con.useCase.CloseSessionsByIds(c, id, &request)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
