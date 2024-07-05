package http

import (
	"authorization/controllers/http/middleware"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	http3 "net/http"
)

func InitServiceMiddleware(handler *gin.Engine, logger logger.ILogger) {
	errorHandler := middleware.NewErrorHandler(logger)
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.Use(errorHandler.HandleError)
	handler.GET("/", func(c *gin.Context) { c.Redirect(http3.StatusPermanentRedirect, "/swagger/index.html") })
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
