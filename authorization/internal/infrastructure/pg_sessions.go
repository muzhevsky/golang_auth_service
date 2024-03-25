package infrastructure

import (
	"authorization/internal/entities"
	"authorization/internal/usecases"
	"authorization/pkg/postgres"
	"context"
	"github.com/Masterminds/squirrel"
)

type sessionRepo struct {
	pg *postgres.Postgres
}

func (s *sessionRepo) Delete(ctx context.Context, session *entities.Session) error {
	sql, args, err := s.pg.Builder.Delete("sessions").
		Where(squirrel.And{squirrel.Eq{"access_token": session.AccessToken}, squirrel.Eq{"refresh_token": session.RefreshToken}}).ToSql()
	if err != nil {
		return err
	}

	_, err = s.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (s *sessionRepo) Create(ctx context.Context, session *entities.Session) error {
	sql, args, err := s.pg.Builder.Insert("sessions").
		Columns("access_token", "refresh_token", "expire_at").
		Values(session.AccessToken, session.RefreshToken, session.ExpireAt).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (s *sessionRepo) Update(ctx context.Context, session *entities.Session) error {
	//TODO implement me
	panic("implement me")
}

func (s *sessionRepo) FindByAccess(ctx context.Context, token string) (*entities.Session, error) {
	sql, args, err := s.pg.Builder.Select("access_token", "refresh_token", "expire_at").
		From("sessions").
		Where(squirrel.Eq{"access_token": token}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	result := &entities.Session{}
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result.AccessToken, &result.RefreshToken, &result.ExpireAt)

	return result, err
}

func NewSessionRepo(pg *postgres.Postgres) usecases.ISessionRepo {
	return &sessionRepo{pg}
}
