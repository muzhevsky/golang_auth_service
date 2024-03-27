package v1

import (
	"authorization/internal/usecases"
	"authorization/pkg/logger"
	"github.com/gin-gonic/gin"
)

func InitServiceMiddleware(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
}

func NewAuthorizationRouter(handler *gin.Engine, u usecases.IUser, l logger.ILogger, s usecases.ISession) {
	h := handler.Group("/auth")

	newAuthRouter(h, u, s, l)
	newSignInRouter(h, u, s, l)
}

func NewAuthenticationRouter(handler *gin.Engine, l logger.ILogger, u usecases.IUser, s usecases.ISession, v usecases.IVerification) {
	h := handler.Group("/user")

	newVerificationRoute(h, v, s, l)
	newSignUpRouter(h, u, v, l)
}
