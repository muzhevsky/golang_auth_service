package user_answers

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

	rows, err := s.client.Pool.Query(context, sql, args...)
	if err != nil {
		return false, nil
	}

	return rows.Next(), nil
}
