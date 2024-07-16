package user_answers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectUserAnswersByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectUserAnswersByAccountIdPGCommand(client *postgres.Client) datasources.ISelectUserAnswersByAccountIdCommand {
	return &selectUserAnswersByAccountIdPGCommand{client: client}
}

func (s *selectUserAnswersByAccountIdPGCommand) Execute(context context.Context, accountId int) (*test_entities.UserTestAnswers, error) {
	sql, args, err := query_builders.NewSelectUserAnswersByAccountIdQuery(&s.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	result := &test_entities.UserTestAnswers{
		AccountId: accountId,
	}
	for rows.Next() {
		row := test_entities.UserTestAnswer{}
		err = rows.Scan(&row.QuestionId, &row.AnswerId)
		if err != nil {
			return nil, err
		}
		result.Answers = append(result.Answers, row)
	}

	return result, nil
}
