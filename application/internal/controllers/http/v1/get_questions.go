package v1

import (
	"github.com/gin-gonic/gin"
	"smartri_app/internal"
	"smartri_app/internal/controllers/http/middleware"
)

type getTestController struct {
	repository internal.ITestRepository
}

func NewGetTestController(repository internal.ITestRepository) *getTestController {
	return &getTestController{repository: repository}
}

func (controller *getTestController) GetQuestions(c *gin.Context) {
	questions, err := controller.repository.GetAllQuestionsWithAnswers(c)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(200, questions)
}
