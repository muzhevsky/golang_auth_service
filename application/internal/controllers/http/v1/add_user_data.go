package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"smartri_app/internal"
	"smartri_app/internal/controllers/http/middleware"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/errs"
	"strconv"
)

type addUserDataController struct {
	useCase internal.IAddUserDataUseCase
}

func NewAddUserDataController(useCase internal.IAddUserDataUseCase) *addUserDataController {
	return &addUserDataController{useCase: useCase}
}

func (controller *addUserDataController) AddUserData(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)
	fmt.Println(err)
	if err != nil {
		middleware.AddGinError(c, errs.SomeErrorToDo)
		return
	}

	var details requests.UserDataRequest
	err = c.ShouldBindJSON(&details)
	fmt.Println(err)
	if err != nil {
		middleware.AddGinError(c, errs.SomeErrorToDo)
		return
	}

	response, err := controller.useCase.Add(c, &details, id)
	fmt.Println(err)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(200, response)
}
