package http

import (
	"authorization/internal/controllers/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitServiceMiddleware(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.Use(middleware.ErrorHandler)
}
