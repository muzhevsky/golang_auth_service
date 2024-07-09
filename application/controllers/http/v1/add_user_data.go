package v1

import (
	"github.com/gin-gonic/gin"
	"smartri_app/controllers/http/middleware"
	"smartri_app/controllers/requests"
	_ "smartri_app/docs"
	"smartri_app/internal"
	"smartri_app/internal/errs"
	"strconv"
)

type addUserDataController struct {
	useCase internal.IInitOrUpdateUserDataUseCase
}

func NewAddUserDataController(useCase internal.IInitOrUpdateUserDataUseCase) *addUserDataController {
	return &addUserDataController{useCase: useCase}
}

// AddUserData godoc
// @Summary      добавляет метаданные пользователя
// @Description  добавление данных пользователя из теста (первые несколько никому не нужных вопросов)
// @Accept       json
// @Produce      json
// @Param request body requests.AddUserDataRequest true "request format"
// @Param Authorization header string true "access token"
// @Success      200  {object} requests.UserDataResponse
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/data [post]
func (controller *addUserDataController) AddUserData(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	var details requests.AddUserDataRequest
	err = c.ShouldBindJSON(&details)
	if err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	response, err := controller.useCase.InitOrUpdate(c, &details, id)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(200, response)
}