package pg

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/pkg/postgres"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type sessionRepo struct {
	pg *postgres.Postgres
}

func NewSessionRepo(pg *postgres.Postgres) datasources.ISessionDataSource {
	return &sessionRepo{pg}
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

func (s *sessionRepo) Create(ctx context.Context, session *entities.Session) (int, error) {
	fmt.Printf("%v ,%v ,%v ,%v ,%v ",
		session.AccessToken, session.RefreshToken, session.UserId, session.DeviceIdentity, session.ExpireAt)
	sql, args, err := s.pg.Builder.Insert("sessions").
		Columns("access_token", "refresh_token", "user_id", "device_identity", "expire_at").
		Values(session.AccessToken, session.RefreshToken, session.UserId, session.DeviceIdentity, session.ExpireAt).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	return id, err
}

func (s *sessionRepo) SelectByAccess(ctx context.Context, token string) (*entities.Session, error) {
	sql, args, err := s.pg.Builder.Select("access_token", "refresh_token", "user_id", "device_identity", "expire_at").
		From("sessions").
		Where(squirrel.Eq{"access_token": token}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	result := &entities.Session{}
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result.AccessToken, &result.RefreshToken, &result.UserId, &result.DeviceIdentity, &result.ExpireAt)

	return result, err
}

func (s *sessionRepo) SelectByUserId(ctx context.Context, userId int) ([]*entities.Session, error) {
	panic("implement me")
	return nil, nil
}
