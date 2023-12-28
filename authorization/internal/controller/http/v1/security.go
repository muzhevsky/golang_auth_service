package v1

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type security struct {
	authUseCase usecase.ISession
	//securityUseCases usecase.ISecurity
	shouldBeAuthorized []string
}

func newSecurity( /*securityUseCases usecase.ISecurity*/ authUseCase usecase.ISession, shouldBeAuthorized []string) *security {
	return &security{authUseCase: authUseCase, shouldBeAuthorized: shouldBeAuthorized}
}

func (handler *security) handle(context *gin.Context) {
	//for _, route := range handler.shouldBeAuthorized {
	//if route == context.Request.URL.Path {
	access := strings.TrimLeft(context.GetHeader("Authorization"), "Bearer")
	access = strings.TrimSpace(access)
	refresh, err := context.Cookie("refresh_token")
	if err != nil {
		println(err)
	}
	refresh = strings.Trim(refresh, " ")
	fmt.Println(handler.checkAuth(context, access, refresh))
	//}
	//}
}

func (handler *security) checkAuth(ctx context.Context, token string, refresh string) bool {
	_, err := handler.authUseCase.VerifyAccessToken(ctx, token)
	if err == nil {
		return true
	}

	if errors.Is(err, entities.TokenExpired) {
		_, err = handler.authUseCase.UpdateAccessToken(ctx, token, refresh)
		if err != nil {
			return false
		}
	}
	return false
}
