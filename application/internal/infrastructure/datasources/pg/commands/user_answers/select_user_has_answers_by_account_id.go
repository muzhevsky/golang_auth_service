package user_answers

import (
	"context"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectUserHasAnswersByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectUserHasAnswersByAccountIdPGCommand(client *postgres.Client) *selectUserHasAnswersByAccountIdPGCommand {
	return &selectUserHasAnswersByAccountIdPGCommand{client: client}
}

func (s *selectUserHasAnswersByAccountIdPGCommand) Execute(context context.Context, accountId int) (bool, error) {
	sql, args, err := query_builders.NewSelectUserAnswerByAccountIdQueryLimit1(&s.client.Builder, accountId)
	if err != nil {
		return false, nil
	}

	rows, err := s.client.Pool.Query(context, sql, args...)
	defer rows.Close()
	if err != nil {
		return false, nil
	}

	return rows.Next(), nil
}
