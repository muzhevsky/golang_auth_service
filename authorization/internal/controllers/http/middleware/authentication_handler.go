package middleware

import (
	"authorization/internal"
	"authorization/internal/errs"
	tokens2 "authorization/internal/infrastructure/services/tokens"
	"github.com/gin-gonic/gin"
	"time"
)

type authenticationHandler struct {
	session internal.ISessionRepository
	manager tokens2.ISessionManager
}

func NewAuthenticationHandler(session internal.ISessionRepository, manager tokens2.ISessionManager) *authenticationHandler {
	return &authenticationHandler{session: session, manager: manager}
}

func (h *authenticationHandler) HandleAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	session, err := h.session.FindByAccessToken(c, token)

	if err != nil {
		c.Set("authError", errs.NotAValidAccessToken)
		c.Next()
		return
	}

	parsedToken, err := h.manager.ParseToken(token)
	if err != nil {
		c.Set("authError", errs.NotAValidAccessToken)
		c.Next()
		return
	}

	expiresAt := time.Unix(int64(parsedToken["expiresAt"].(float64)), 0)

	if expiresAt.Before(time.Now()) {
		c.Set("authError", errs.AccessTokenExpired)
		c.Next()
		return
	}

	c.Set("userId", session.UserId)
	c.Next()
}
