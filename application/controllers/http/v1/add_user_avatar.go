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

type initOrUpdateUserAvatarController struct {
	useCase internal.IInitOrUpdateAvatarUseCase
}

func NewInitOrUpdateUserAvatarController(useCase internal.IInitOrUpdateAvatarUseCase) *initOrUpdateUserAvatarController {
	return &initOrUpdateUserAvatarController{useCase: useCase}
}

// InitOrUpdateAvatar godoc
// @Summary      обновляет аватар пользователя
// @Description  обновляет аватар пользователя или создает его, если тот ещё не создан access token
// @Accept       json
// @Produce      json
// @Param request body requests.AvatarRequest true "request format"
// @Param Authorization header string true "access token"
// @Success      200  {object} requests.AvatarRequest
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/avatar [post]
func (controller *initOrUpdateUserAvatarController) InitOrUpdateAvatar(c *gin.Context) {
	accountId := c.GetHeader("account_id")
	id, err := strconv.Atoi(accountId)
	if err != nil {
		middleware.AddGinError(c, errs.UnauthenticatedError)
		return
	}

	request := requests.AvatarRequest{}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		middleware.AddGinError(c, errs.DataBindError)
		return
	}

	err = controller.useCase.InitOrUpdate(c, id, &request)
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
