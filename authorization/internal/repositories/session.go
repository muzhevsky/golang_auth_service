package repositories

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type sessionRepository struct {
	ds datasources.ISessionDataSource
}

func NewSessionRepository(ds datasources.ISessionDataSource) *sessionRepository {
	return &sessionRepository{ds: ds}
}

func (s *sessionRepository) Create(context context.Context, session *entities.Session) (int, error) {
	return s.ds.Create(context, session)
}

func (s *sessionRepository) FindByAccessToken(context context.Context, token string) (*entities.Session, error) {
	return s.ds.SelectByAccess(context, token)
}
