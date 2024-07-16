package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartri_app/controllers/http/middleware"
	"smartri_app/controllers/requests"
	_ "smartri_app/docs"
	"smartri_app/internal"
	"smartri_app/internal/errs"
	"strconv"
)

type addUserAnswersController struct {
	useCase internal.IAddUserTestAnswersUseCase
}

func NewAddUserAnswersController(useCase internal.IAddUserTestAnswersUseCase) *addUserAnswersController {
	return &addUserAnswersController{useCase: useCase}
}

// AddUserAnswers godoc
// @Summary      добавление ответов на тест
// @Description  добавляет пользовательские ответы на входное тестирование
// @Accept       json
// @Produce      json
// @Param request body requests.UserAnswersRequest true "request format"
// @Success      200  "ok"
// @Param Authorization header string true "access token"
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 404 {object} middleware.ErrorResponse "не найдены данные пользователя, сначала нужно отправить их /user/data [post]"
// @Failure 409 {object} middleware.ErrorResponse "пользователь уже прошел тест ранее"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/test_entities [post]
func (controller *addUserAnswersController) AddUserAnswers(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)

	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	var answers requests.UserAnswersRequest
	err = c.ShouldBindJSON(&answers)

	if err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	response, err := controller.useCase.Add(c, &answers, id)

	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
