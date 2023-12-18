package v1

import (
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServiceMiddleware(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)
}

func NewAuthorizationRouter(handler *gin.Engine, u usecase.IUser, l logger.ILogger, s usecase.ISession) {
	h := handler.Group("/auth")

	newSignInRouter(h, u, s, l)
}

func NewAuthenticationRouter(handler *gin.Engine, l logger.ILogger, u usecase.IUser, s usecase.ISession, v usecase.IVerification) {
	h := handler.Group("/user")

	newVerificationRoute(h, v, s, l)
	newSignUpRouter(h, u, v, l)

}
