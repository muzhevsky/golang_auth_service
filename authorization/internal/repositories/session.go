package repositories

import (
	"authorization/internal"
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type sessionRepository struct {
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand
	selectSessionsByAccountIdCommand  datasources.ISelectSessionsByAccountIdCommand
	insertSessionCommand              datasources.IInsertSessionCommand
	updateSessionByAccessTokenCommand datasources.IUpdateSessionByAccessTokenCommand
}

func NewSessionRepository(
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand,
	selectSessionsByAccountIdCommand datasources.ISelectSessionsByAccountIdCommand,
	insertSessionCommand datasources.IInsertSessionCommand,
	updateSessionByAccessTokenCommand datasources.IUpdateSessionByAccessTokenCommand) internal.ISessionRepository {
	return &sessionRepository{selectSessionByAccessTokenCommand: selectSessionByAccessTokenCommand, selectSessionsByAccountIdCommand: selectSessionsByAccountIdCommand, insertSessionCommand: insertSessionCommand, updateSessionByAccessTokenCommand: updateSessionByAccessTokenCommand}
}

func (s *sessionRepository) Create(context context.Context, session *session.Session) error {
	return s.insertSessionCommand.Execute(context, session)
}

func (s *sessionRepository) FindByAccessToken(context context.Context, token string) (*session.Session, error) {
	return s.selectSessionByAccessTokenCommand.Execute(context, token)
}

func (s *sessionRepository) Update(context context.Context, sessionToUpdate *session.Session, newSession *session.Session) (*session.Session, error) {
	cpy := session.Session{
		AccountId:    newSession.AccountId,
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
		ExpiresAt:    newSession.ExpiresAt,
	}
	err := s.updateSessionByAccessTokenCommand.Execute(context, sessionToUpdate.AccessToken, &cpy)

	return sessionToUpdate, err
}
