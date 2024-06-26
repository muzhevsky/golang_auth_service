package repositories

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type sessionRepository struct {
	ds datasources.ISessionDatasource
}

func NewSessionRepository(ds datasources.ISessionDatasource) *sessionRepository {
	return &sessionRepository{ds: ds}
}

func (s *sessionRepository) Create(context context.Context, session *entities.Session) (int, error) {
	return s.ds.Create(context, session)
}

func (s *sessionRepository) FindByAccessToken(context context.Context, token string) (*entities.Session, error) {
	return s.ds.SelectByAccess(context, token)
}

func (s *sessionRepository) Update(context context.Context, sessionToUpdate *entities.Session, newSession *entities.Session) (*entities.Session, error) {
	id := sessionToUpdate.Id

	sessionToUpdate.ExpiresAt = newSession.ExpiresAt
	sessionToUpdate.AccessToken = newSession.AccessToken
	sessionToUpdate.RefreshToken = newSession.RefreshToken

	err := s.ds.UpdateById(context, id, sessionToUpdate)

	return sessionToUpdate, err
}
