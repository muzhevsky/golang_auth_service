package repositories

import (
	"authorization/internal"
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type sessionRepository struct {
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand
	insertSessionCommand              datasources.IInsertSessionCommand
	updateSessionByAccessTokenCommand datasources.IUpdateSessionByAccessTokenCommand
	deleteSessionByAccessTokenCommand datasources.IDeleteSessionByAccessTokenCommand
}

func NewSessionRepository(
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand,
	insertSessionCommand datasources.IInsertSessionCommand,
	updateSessionByAccessTokenCommand datasources.IUpdateSessionByAccessTokenCommand,
	deleteSessionByAccessTokenCommand datasources.IDeleteSessionByAccessTokenCommand) internal.ISessionRepository {

	return &sessionRepository{
		selectSessionByAccessTokenCommand: selectSessionByAccessTokenCommand,
		insertSessionCommand:              insertSessionCommand,
		updateSessionByAccessTokenCommand: updateSessionByAccessTokenCommand,
		deleteSessionByAccessTokenCommand: deleteSessionByAccessTokenCommand}
}

func (s *sessionRepository) Create(context context.Context, se *session.Session) error {
	return s.insertSessionCommand.Execute(context, se)
}

func (s *sessionRepository) FindByAccessToken(context context.Context, token string) (*session.Session, error) {
	return s.selectSessionByAccessTokenCommand.Execute(context, token)
}

func (s *sessionRepository) UpdateByAccessToken(context context.Context, token string, newSession *session.Session) (*session.Session, error) {
	cpy := session.Session{
		AccountId:    newSession.AccountId,
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
		ExpiresAt:    newSession.ExpiresAt,
	}

	err := s.updateSessionByAccessTokenCommand.Execute(context, token, &cpy)
	if err != nil {
		return nil, err
	}

	return newSession, err
}

func (s *sessionRepository) DeleteByAccessToken(context context.Context, token string) error {
	return s.deleteSessionByAccessTokenCommand.Execute(context, token)
}
