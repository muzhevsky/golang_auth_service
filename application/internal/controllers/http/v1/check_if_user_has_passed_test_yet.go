package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smartri_app/internal"
	"smartri_app/internal/controllers/http/middleware"
	"smartri_app/internal/errs"
	"strconv"
)

type checkIfUserHasPassedTestYetController struct {
	useCase internal.ICheckIfUserHasPassedTestYetUseCase
}

func NewCheckIfUserHasPassedTestYetController(useCase internal.ICheckIfUserHasPassedTestYetUseCase) *checkIfUserHasPassedTestYetController {
	return &checkIfUserHasPassedTestYetController{useCase: useCase}
}

func (controller *checkIfUserHasPassedTestYetController) Check(c *gin.Context) {
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
