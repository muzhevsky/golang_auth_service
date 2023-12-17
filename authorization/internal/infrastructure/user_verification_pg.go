package infrastructure

import (
	"authorization/internal/entities"
	"authorization/internal/usecase"
	"authorization/pkg/postgres"
	"context"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type verificationRepo struct {
	pg *postgres.Postgres
}

func NewVerificationRepo(pg *postgres.Postgres) usecase.IVerificationRepo {
	return &verificationRepo{pg}
}

func (repo *verificationRepo) Create(verification *entities.Verification) error {
	request := "select add_verification_code($1,$2,$3)"

	_, err := repo.pg.Pool.Exec(context.Background(), request, verification.UserId, verification.Code, verification.ExpiredTime)
	if err != nil {
		return err
	}
	return nil
}

func (repo *verificationRepo) FindOne(userId int) (*entities.Verification, error) {
	request, args, err := repo.pg.Builder.Select("verification_code", "expiration_time").
		From("verification_codes").Where(sq.Eq{"user_id": userId}).Limit(1).ToSql()

	rows, err := repo.pg.Pool.Query(context.Background(), request, args...)
	if err != nil {
		return nil, err
	}

	var verificationCode string
	var expirationTime time.Time
	for rows.Next() {
		err = rows.Scan(&verificationCode, &expirationTime)
		if err != nil {
			return nil, err
		}
	}

	return &entities.Verification{userId, verificationCode, expirationTime}, nil
}
