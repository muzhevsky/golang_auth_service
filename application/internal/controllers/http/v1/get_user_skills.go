package v1

import (
	"github.com/gin-gonic/gin"
	_ "smartri_app/docs"
	"smartri_app/internal"
	"smartri_app/internal/controllers/http/middleware"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/errs"
	"strconv"
)

type getUserSkillDataController struct {
	repository internal.ISkillRepository
}

func NewGetSkillDataController(repository internal.ISkillRepository) *getUserSkillDataController {
	return &getUserSkillDataController{repository: repository}
}

// GetUserSkills godoc
// @Summary      получает данные о скиллах пользователя
// @Description  получает данные о скиллах пользователя используя access token
// @Produce      json
// @Param Authorization header string true "access token"
// @Success      200  {object} []entities.UserSkills
// @Failure 401 {object} middleware.ErrorResponse "некорректный access token"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/skills [get]
func (controller *getUserSkillDataController) GetUserSkills(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	skills, err := controller.repository.GetSkillsByAccountId(c, id)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(200, requests.UserSkillResponse{
		AccountId: id,
		Skills:    skills,
	})

}
