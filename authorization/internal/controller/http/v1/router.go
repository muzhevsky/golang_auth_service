package v1

import (
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine, l logger.ILogger, u usecase.IUser, v usecase.IVerification, s usecase.ISession) { // todo явно надо роутеры по-другому создавать
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	h := handler.Group("/user")
	{
		newUserRoutes(h, u, v, l)
	}
	h = handler.Group("/session")
	{
		newSessionRoutes(h, s, l)
	}
}
