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

type addUserXpController struct {
	uc internal.IAddUserXpChangeUseCase
}

func NewAddUserXpController(uc internal.IAddUserXpChangeUseCase) *addUserXpController {
	return &addUserXpController{uc: uc}
}

// AddUserXp godoc
// @Summary      добавляет опыт пользователю
// @Description  добавляет опыт пользователю по скиллам используя access token
// @Accept       json
// @Produce      json
// @Param request body requests.AddSkillChangeRequest true "request format"
// @Param Authorization header string true "access token"
// @Success      200  {object} requests.UserDataResponse
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/xp [post]
func (controller *addUserXpController) AddUserXp(c *gin.Context) {
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
