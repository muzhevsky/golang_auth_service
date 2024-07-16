package v1

import (
	"github.com/gin-gonic/gin"
	"smartri_app/controllers/http/middleware"
	_ "smartri_app/docs"
	"smartri_app/internal"
)

type getTestController struct {
	repository internal.ITestRepository
}

func NewGetTestController(repository internal.ITestRepository) *getTestController {
	return &getTestController{repository: repository}
}

// GetQuestions godoc
// @Summary      получает вопросы для входного теста
// @Description  получает вопросы для входного теста
// @Accept       json
// @Produce      json
// @Success      200  {object} []test_entities.Question
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /test_entities [get]
func (controller *getTestController) GetQuestions(c *gin.Context) {
	questions, err := controller.repository.GetAllQuestionsWithAnswers(c)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}
	c.JSON(200, questions)
}
