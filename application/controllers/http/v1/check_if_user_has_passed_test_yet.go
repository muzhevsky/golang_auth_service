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

type checkIfUserHasPassedTestYetController struct {
	useCase internal.ICheckIfUserHasPassedTestYetUseCase
}

func NewCheckIfUserHasPassedTestYetController(useCase internal.ICheckIfUserHasPassedTestYetUseCase) *checkIfUserHasPassedTestYetController {
	return &checkIfUserHasPassedTestYetController{useCase: useCase}
}

// CheckIfUserHasPassedTestYet godoc
// @Summary      проверка на прохождение теста
// @Description  прошел ли пользователь тест
// @Param Authorization header string true "access token"
// @Success      200
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /test_entities/passed [get]
func (controller *checkIfUserHasPassedTestYetController) CheckIfUserHasPassedTestYet(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	passed, err := controller.useCase.Check(c, id)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, passed)
}
