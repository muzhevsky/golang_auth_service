package v1

import (
	http2 "authorization/internal/controllers/http/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authenticationController struct {
}

func NewAuthenticationController(handler *gin.Engine) {
	u := &authenticationController{}

	handler.GET("/authenticate", u.authenticate)
}

func (r *authenticationController) authenticate(c *gin.Context) {
	accountId, exists := c.Get("accountId")
	if !exists {
		err, _ := c.Get("authError")
		http2.AddGinError(c, err.(error))
		return
	}

	c.JSON(http.StatusOK, accountId)
}
