package user_data

import (
	"context"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectUserHasAnswersByAccountIdCommand struct {
	client *postgres.Client
}

func NewSelectUserHasAnswersByAccountIdCommand(client *postgres.Client) *selectUserHasAnswersByAccountIdCommand {
	return &selectUserHasAnswersByAccountIdCommand{client: client}
}

func (s *selectUserHasAnswersByAccountIdCommand) Execute(context context.Context, accountId int) (bool, error) {
	sql, args, err := query_builders.NewSelectUserAnswerByAccountIdQueryLimit1(&s.client.Builder, accountId)
	if err != nil {
		return false, nil
	}

	row := s.client.Pool.QueryRow(context, sql, args...)

	err = row.Scan()
	if err != nil {
		return false, nil
	}
	return true, nil
}
