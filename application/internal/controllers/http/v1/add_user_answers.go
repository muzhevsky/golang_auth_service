package v1

import (
	"github.com/gin-gonic/gin"
	"smartri_app/internal"
	"smartri_app/internal/controllers/http/middleware"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/errs"
	"strconv"
)

type addUserAnswersController struct {
	useCase internal.IAddUserTestAnswersUseCase
}

func NewAddUserAnswersController(useCase internal.IAddUserTestAnswersUseCase) *addUserAnswersController {
	return &addUserAnswersController{useCase: useCase}
}

func (controller *addUserAnswersController) AddUserAnswers(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)

	if err != nil {
		middleware.AddGinError(c, errs.SomeErrorToDo)
		return
	}

	var answers requests.UserAnswersRequest
	err = c.ShouldBindJSON(&answers)

	if err != nil {
		middleware.AddGinError(c, errs.SomeErrorToDo)
		return
	}

	err = controller.useCase.Add(c, &answers, id)

	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(200, "ok")
}
