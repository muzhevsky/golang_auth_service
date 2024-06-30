package questions

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAllQuestionsCommand struct {
	client *postgres.Client
}

func NewSelectAllQuestionsCommand(client *postgres.Client) datasources.ISelectAllQuestionsCommand {
	return &selectAllQuestionsCommand{client: client}
}

func (q *selectAllQuestionsCommand) Execute(context context.Context) ([]*entities.Question, error) {
	sql, args, err := query_builders.NewSelectAllQuestionsQuery(&q.client.Builder)
	if err != nil {
		return nil, err
	}

	rows, err := q.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*entities.Question, 0)
	for rows.Next() {
		question := entities.Question{}
		err = rows.Scan(&question.Id, &question.Text)
		if err != nil {
			return nil, err
		}
		result = append(result, &question)
	}

	return result, nil
}
