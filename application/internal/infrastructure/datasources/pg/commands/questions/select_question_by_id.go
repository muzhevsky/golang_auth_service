package questions

import (
	"context"
	"smartri_app/internal/entities/test"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectQuestionByIdPGCommand struct {
	client *postgres.Client
}

func NewSelectQuestionByIdPGCommand(client *postgres.Client) *selectQuestionByIdPGCommand {
	return &selectQuestionByIdPGCommand{client: client}
}

func (s *selectQuestionByIdPGCommand) Execute(context context.Context, id int) (*test.Question, error) {
	sql, args, err := query_builders.NewSelectQuestionByIdQuery(&s.client.Builder, id)
	if err != nil {
		return nil, err
	}

	row := s.client.Pool.QueryRow(context, sql, args...)

	var result test.Question
	err = row.Scan(&result.Id, &result.Text)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
