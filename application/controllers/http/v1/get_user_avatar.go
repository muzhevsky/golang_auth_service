package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartri_app/controllers/http/middleware"
	_ "smartri_app/docs"
	"smartri_app/internal"
	"smartri_app/internal/errs"
	"strconv"
)

type getUserAvatarController struct {
	useCase internal.IGetUserAvatarUseCase
}

func NewGetUserAvatarController(useCase internal.IGetUserAvatarUseCase) *getUserAvatarController {
	return &getUserAvatarController{useCase: useCase}
}

// GetUserAvatar godoc
// @Summary      получает аватар пользователя
// @Description  получает аватар пользователя по access token
// @Accept       json
// @Produce      json
// @Param Authorization header string true "access token"
// @Success      200  {object} requests.AvatarRequest
// @Failure 400 {object} middleware.ErrorResponse "ошибка формата отправленных клиентом данных"
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/avatar [get]
func (controller *getUserAvatarController) GetUserAvatar(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	response, err := controller.useCase.GetAvatar(c, id)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
