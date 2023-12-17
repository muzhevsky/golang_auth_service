package infrastructure

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"authorization/pkg/postgres"
	"context"
	"github.com/Masterminds/squirrel"
)

type sessionRepo struct {
	pg *postgres.Postgres
}

func (s sessionRepo) Create(ctx context.Context, session *entities.Session) error {
	sql, args, err := s.pg.Builder.Insert("sessions").
		Columns("user_id", "access_token", "refresh_token", "expire_at").
		Values(session.Id, session.AccessToken, session.RefreshToken, session.ExpireAt).
		ToSql()
	if err != nil {
		return err
	}

	err = s.pg.Pool.QueryRow(ctx, sql, args).Scan(&squirrel.Row{})
	return err
}

func (s sessionRepo) Update(ctx context.Context, session *entities.Session) error {
	//TODO implement me
	panic("implement me")
}

func (s sessionRepo) FindByAccess(ctx context.Context, token string) (*entities.Session, error) {
	sql, args, err := s.pg.Builder.Select("user_id", "access_token", "refresh_token", "expire_at").
		From("users").
		Where(squirrel.Eq{"access_token": token}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	var result *entities.Session
	err = s.pg.Pool.QueryRow(ctx, sql, args).Scan(result)

	return result, err
}

func NewSessionRepo(pg *postgres.Postgres) usecase.ISessionRepo {
	return &sessionRepo{pg}
}
