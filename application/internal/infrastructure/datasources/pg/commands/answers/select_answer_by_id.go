package answers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/test"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAnswerByIdPGCommand struct {
	client *postgres.Client
}

func NewSelectAnswerByIdPGCommand(client *postgres.Client) datasources.ISelectAnswerByIdCommand {
	return &selectAnswerByIdPGCommand{client: client}
}

func (s *selectAnswerByIdPGCommand) Execute(context context.Context, id int) (*test.Answer, error) {
	sql, args, err := query_builders.NewSelectAnswerByIdQuery(&s.client.Builder, id)
	if err != nil {
		return nil, err
	}

	row := s.client.Pool.QueryRow(context, sql, args...)
	var answer test.Answer

	answer.Id = id
	err = row.Scan(&answer.Text, &answer.QuestionId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &answer, nil
}
