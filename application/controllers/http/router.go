package http

import (
	"github.com/gin-gonic/gin"
	"smartri_app/controllers/http/middleware"
	"smartri_app/pkg/logger"
)

func InitServiceMiddleware(handler *gin.Engine, logger logger.ILogger) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	errorHandler := middleware.NewErrorHandler(logger)
	handler.Use(errorHandler.HandleError)
}
