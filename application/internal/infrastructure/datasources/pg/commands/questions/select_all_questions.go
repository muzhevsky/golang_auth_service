package questions

import (
	"context"
	"smartri_app/internal/entities/test"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAllQuestionsPGCommand struct {
	client *postgres.Client
}

func NewSelectAllQuestionsPGCommand(client *postgres.Client) datasources.ISelectAllQuestionsCommand {
	return &selectAllQuestionsPGCommand{client: client}
}

func (q *selectAllQuestionsPGCommand) Execute(context context.Context) ([]*test.Question, error) {
	sql, args, err := query_builders.NewSelectAllQuestionsQuery(&q.client.Builder)
	if err != nil {
		return nil, err
	}

	rows, err := q.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*test.Question, 0)
	for rows.Next() {
		question := test.Question{}
		err = rows.Scan(&question.Id, &question.Text)
		if err != nil {
			return nil, err
		}
		result = append(result, &question)
	}

	return result, nil
}
