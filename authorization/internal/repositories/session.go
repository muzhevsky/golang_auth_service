package repositories

import (
	"authorization/internal"
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"context"
	"fmt"
)

type sessionRepository struct {
	selectSessionByIdCommand          datasources.ISelectSessionByIdCommand
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand
	selectSessionsByAccountIdCommand  datasources.ISelectSessionsByAccountIdCommand
	insertSessionCommand              datasources.IInsertSessionCommand
	updateSessionByIdCommand          datasources.IUpdateSessionByIdCommand
}

func NewSessionRepository(
	selectSessionByIdCommand datasources.ISelectSessionByIdCommand,
	selectSessionByAccessTokenCommand datasources.ISelectSessionByAccessTokenCommand,
	selectSessionsByAccountIdCommand datasources.ISelectSessionsByAccountIdCommand,
	insertSessionCommand datasources.IInsertSessionCommand,
	updateSessionByIdCommand datasources.IUpdateSessionByIdCommand) internal.ISessionRepository {
	return &sessionRepository{selectSessionByIdCommand: selectSessionByIdCommand, selectSessionByAccessTokenCommand: selectSessionByAccessTokenCommand, selectSessionsByAccountIdCommand: selectSessionsByAccountIdCommand, insertSessionCommand: insertSessionCommand, updateSessionByIdCommand: updateSessionByIdCommand}
}

func (s *sessionRepository) Create(context context.Context, session *entities.Session) (int, error) {
	fmt.Println("a")
	return s.insertSessionCommand.Execute(context, session)
}

func (s *sessionRepository) FindByAccessToken(context context.Context, token string) (*entities.Session, error) {
	return s.selectSessionByAccessTokenCommand.Execute(context, token)
}

func (s *sessionRepository) Update(context context.Context, sessionToUpdate *entities.Session, newSession *entities.Session) (*entities.Session, error) {
	cpy := entities.Session{
		Id:           sessionToUpdate.Id,
		AccountId:    newSession.AccountId,
		AccessToken:  newSession.AccessToken,
		RefreshToken: newSession.RefreshToken,
		ExpiresAt:    newSession.ExpiresAt,
	}
	err := s.updateSessionByIdCommand.Execute(context, &cpy)

	return sessionToUpdate, err
}
