package questions

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectQuestionByIdCommand struct {
	client *postgres.Client
}

func NewSelectQuestionByIdCommand(client *postgres.Client) *selectQuestionByIdCommand {
	return &selectQuestionByIdCommand{client: client}
}

func (s *selectQuestionByIdCommand) Execute(context context.Context, id int) (*entities.Question, error) {
	sql, args, err := query_builders.NewSelectQuestionByIdQuery(&s.client.Builder, id)
	if err != nil {
		return nil, err
	}

	row := s.client.Pool.QueryRow(context, sql, args...)

	var result entities.Question
	err = row.Scan(&result.Id, &result.Text)
	if err != nil {
		return nil, err
	}

	return &result, nil

}
