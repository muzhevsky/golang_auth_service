package infrastructure

import "authorization/pkg/postgres"

type UserDataRepo struct {
	*postgres.Postgres
}

func NewUserDataRepo(pg *postgres.Postgres) *UserDataRepo {
	return &UserDataRepo{pg}
}

// TODO implement
