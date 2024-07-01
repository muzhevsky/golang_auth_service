package v1

import (
	"github.com/gin-gonic/gin"
	"smartri_app/internal"
	"smartri_app/internal/controllers/http/middleware"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/errs"
	"strconv"
)

type addUserXpController struct {
	uc internal.IAddUserXpChangeUseCase
}

func NewAddUserXpController(uc internal.IAddUserXpChangeUseCase) *addUserXpController {
	return &addUserXpController{uc: uc}
}

func (controller *addUserXpController) Add(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	request := requests.AddSkillChangeRequest{}

	err = c.ShouldBindJSON(&request)
	if err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	skills, err := controller.uc.Add(c, id, &request)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(200, skills)
}
