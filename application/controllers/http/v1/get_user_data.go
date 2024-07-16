package v1

import (
	"github.com/gin-gonic/gin"
	"smartri_app/controllers/http/middleware"
	"smartri_app/internal"
	"smartri_app/internal/errs"
	"strconv"
)

type getUserDataController struct {
	repo internal.IUserDataRepository
}

func NewGetUserDataController(repo internal.IUserDataRepository) *getUserDataController {
	return &getUserDataController{repo: repo}
}

// GetUserData godoc
// @Summary      получает метаданные пользователя
// @Description  получение метаданных пользователя (пол, возраст, ник, общий опыт) с использованием access_token в
// заголовке Authorization
// @Accept       json
// @Produce      json
// @Param Authorization header string true "access token"
// @Success      200  {object} requests.UserDataResponse
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/data [get]
func (controller *getUserDataController) GetUserData(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)

	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	userData, err := controller.repo.GetByAccountId(c, id)

	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(200, userData)
}
