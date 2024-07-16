package v1

import (
	"authorization/controllers/http/middleware"
	_ "authorization/docs"
	"authorization/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type requestVerificationController struct {
	user         internal.ICreateAccountUseCase
	verification internal.IRequestVerificationUseCase
}

func NewRequestVerificationController(handler *gin.Engine, user internal.ICreateAccountUseCase, verification internal.IRequestVerificationUseCase) {
	u := &requestVerificationController{user, verification}

	handler.POST("/verification/request", u.requestVerification)
}

// RequestVerification godoc
// @Summary      запрос на верификацию пользователя
// @Description запрос на верификацию пользователя с использованием токена, переданного в заголовке "Authorization"
// @Accept       json
// @Produce      json
// @Success      200  "(TODO: в данные момент возвращается код, чтобы каждый раз клиент не мучал почту) Ok"
// @Param Authorization header string true "access token"
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "некорректный access token"
// @Failure 409 {object} middleware.ErrorResponse "пользователь уже верифицирован"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /verification/request [post]
func (u *requestVerificationController) requestVerification(c *gin.Context) {
	accountId, exists := c.Get("accountId")
	if !exists {
		err, _ := c.Get("authError")
		middleware.AddGinError(c, err.(error))
		return
	}

	code, err := u.verification.RequestVerification(c, accountId.(int))
	if err != nil {
		middleware.AddGinError(c, err)
		return
	}

	c.JSON(http.StatusOK, code)
}
