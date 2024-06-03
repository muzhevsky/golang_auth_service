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

const sessionsTableName = "sessions"

func NewSessionRepo(pg *postgres.Postgres) datasources.ISessionDataSource {
	return &sessionRepo{pg}
}

func (s *sessionRepo) Delete(ctx context.Context, session *entities.Session) error {
	sql, args, err := s.pg.Builder.Delete(sessionsTableName).
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
	sql, args, err := s.pg.Builder.Insert(sessionsTableName).
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
		From(sessionsTableName).
		Where(squirrel.Eq{"access_token": token}).
		Limit(1).
		ToSql()

	fmt.Println(token)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	result := &entities.Session{}
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result.AccessToken, &result.RefreshToken, &result.UserId, &result.DeviceIdentity, &result.ExpireAt)

	fmt.Println(result.AccessToken, err)
	return result, err
}

func (s *sessionRepo) UpdateById(ctx context.Context, id int, session *entities.Session) error {
	sql, args, err := s.pg.Builder.Update(sessionsTableName).
		Set("access_token", session.AccessToken).
		Set("refresh_token", session.RefreshToken).
		Set("user_id", session.UserId).
		Set("device_identity", session.DeviceIdentity).
		Set("expire_at", session.ExpireAt).ToSql()
	if err != nil {
		return err
	}

	_, err = s.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (s *sessionRepo) SelectByUserId(ctx context.Context, userId int) ([]*entities.Session, error) {
	panic("implement me")
	return nil, nil
}
