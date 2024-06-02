package v1

//
//import (
//	"authorization/internal"
//	"authorization/internal/entities"
//	"authorization/pkg/logger"
//	"errors"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"time"
//)
//
//type authRouter struct {
//	user internal.IUser
//	auth internal.ISession
//	logger    logger.ILogger
//}
//
//func newAuthRouter(handler *gin.RouterGroup, user internal.IUser, auth internal.ISession, logger logger.ILogger) {
//	u := &authRouter{user, auth, logger}
//
//	handler.POST("/auth", u.authorize)
//}
//
//type authRequest struct {
//	AccessToken  string `json:"accessToken"`
//	RefreshToken string `json:"refreshToken"`
//}
//
//type authResponse struct {
//	UserId       int    `json:"userId"`
//	AccessToken  string `json:"accessToken"`
//	RefreshToken string `json:"refreshToken"`
//}
//
//func (r *authRouter) authorize(c *gin.Context) {
//	request := authRequest{}
//	if err := c.ShouldBind(&request); err != nil {
//		r.logger.Error(err, "http - v1 - verifyUser")
//		errorResponse(c, http.StatusBadRequest, "invalid request body", DataBindErrorCode)
//		return
//	}
//	response := authResponse{}
//	err := r.auth.VerifyAccessToken(c, request.AccessToken)
//	if err != nil {
//		if errors.Is(err, entities.AccessTokenExpired) {
//			session := entities.Session{
//				AccessToken:  request.AccessToken,
//				RefreshToken: request.RefreshToken,
//				ExpireAt:     time.Now(),
//			}
//			updatedSession, updateErr := r.auth.UpdateSession(c, &session)
//			if updateErr == nil {
//				response.AccessToken = updatedSession.AccessToken
//				response.RefreshToken = updatedSession.RefreshToken
//				response.UserId = updatedSession.UserId
//				c.JSON(http.StatusOK, response)
//			}
//			err = updateErr
//		}
//		if err != nil {
//			if errors.Is(err, entities.RefreshTokenExpired) ||
//				errors.Is(err, entities.NotAValidAccessToken) ||
//				errors.Is(err, entities.NotAValidRefreshToken) {
//				errorResponse(c, http.StatusUnauthorized, err.Error(), InternalServerErrorErrorCode)
//				return
//			}
//
//			r.logger.Error("", err.Error())
//			errorResponse(c, http.StatusUnauthorized, "unexpected error", InternalServerErrorErrorCode)
//			return
//		}
//		return
//	}
//
//	session, _ := r.auth.GetSession(c, request.AccessToken)
//	response.AccessToken = session.AccessToken
//	response.RefreshToken = session.RefreshToken
//	response.UserId = session.UserId
//	c.JSON(http.StatusOK, response)
//}
