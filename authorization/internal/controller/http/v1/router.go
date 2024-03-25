package v1

import (
	"authorization/internal/usecase"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
)

func InitServiceMiddleware(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
}

func NewAuthorizationRouter(handler *gin.Engine, u usecase.IUser, l logger.ILogger, s usecase.ISession) {
	h := handler.Group("/auth")

	newAuthRouter(h, u, s, l)
	newSignInRouter(h, u, s, l)
}

func NewAuthenticationRouter(handler *gin.Engine, l logger.ILogger, u usecase.IUser, s usecase.ISession, v usecase.IVerification) {
	h := handler.Group("/user")

	newVerificationRoute(h, v, s, l)
	newSignUpRouter(h, u, v, l)
}
