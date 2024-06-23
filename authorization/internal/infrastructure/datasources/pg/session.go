package pg

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/pkg/postgres"
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
)

type sessionDatasource struct {
	pg *postgres.Client
}

const sessionsTableName = "sessions"

func NewSessionDatasource(pg *postgres.Client) datasources.ISessionDatasource {
	return &sessionDatasource{pg}
}

func (s *sessionDatasource) Delete(ctx context.Context, session *entities.Session) error {
	sql, args, err := s.pg.Builder.Delete(sessionsTableName).
		Where(squirrel.And{squirrel.Eq{"access_token": session.AccessToken}, squirrel.Eq{"refresh_token": session.RefreshToken}}).ToSql()
	if err != nil {
		return err
	}

	_, err = s.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (s *sessionDatasource) Create(ctx context.Context, session *entities.Session) (int, error) {
	sql, args, err := s.pg.Builder.Insert(sessionsTableName).
		Columns("access_token", "refresh_token", "user_id", "device_identity", "expire_at").
		Values(session.AccessToken, session.RefreshToken, session.AccountId, session.DeviceIdentity, session.ExpiresAt).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	return id, err
}

func (s *sessionDatasource) SelectByAccess(ctx context.Context, token string) (*entities.Session, error) {
	sql, args, err := s.pg.Builder.Select("access_token", "refresh_token", "user_id", "device_identity", "expire_at").
		From(sessionsTableName).
		Where(squirrel.Eq{"access_token": token}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	result := &entities.Session{}
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result.AccessToken, &result.RefreshToken, &result.AccountId, &result.DeviceIdentity, &result.ExpiresAt)

	if errors.Is(err, s.pg.ErrNoRows) {
		return nil, nil
	}
	return result, err
}

func (s *sessionDatasource) UpdateById(ctx context.Context, id int, session *entities.Session) error {
	sql, args, err := s.pg.Builder.Update(sessionsTableName).
		Set("access_token", session.AccessToken).
		Set("refresh_token", session.RefreshToken).
		Set("user_id", session.AccountId).
		Set("device_identity", session.DeviceIdentity).
		Set("expire_at", session.ExpiresAt).ToSql()
	if err != nil {
		return err
	}

	_, err = s.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (s *sessionDatasource) SelectByUserId(ctx context.Context, userId int) ([]*entities.Session, error) {
	panic("implement me")
	return nil, nil
}
