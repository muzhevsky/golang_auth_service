package user_answers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectUserAnswersByAccountIdCommand struct {
	client *postgres.Client
}

func NewSelectUserAnswersByAccountIdCommand(client *postgres.Client) datasources.ISelectUserAnswersByAccountIdCommand {
	return &selectUserAnswersByAccountIdCommand{client: client}
}

func (s *selectUserAnswersByAccountIdCommand) Execute(context context.Context, accountId int) (*entities.UserTestAnswers, error) {
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

	result := &entities.UserTestAnswers{
		AccountId: accountId,
	}
	for rows.Next() {
		row := entities.UserTestAnswer{}
		err = rows.Scan(&row.QuestionId, &row.AnswerId)
		if err != nil {
			return nil, err
		}
		result.Answers = append(result.Answers, row)
	}

	return result, nil
}
