package answers

import (
	"context"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAnswersByQuestionIdPGCommand struct {
	client *postgres.Client
}

func NewSelectAnswersByQuestionIdPGCommand(client *postgres.Client) *selectAnswersByQuestionIdPGCommand {
	return &selectAnswersByQuestionIdPGCommand{client: client}
}

func (q *selectAnswersByQuestionIdPGCommand) Execute(context context.Context, questionId int) ([]*test_entities.Answer, error) {
	sql, args, err := query_builders.NewSelectAnswersByQuestionIdQuery(&q.client.Builder, questionId)
	if err != nil {
		return nil, err
	}

	rows, err := q.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*test_entities.Answer, 0)
	for rows.Next() {
		answer := test_entities.Answer{}
		err = rows.Scan(&answer.Id, &answer.Text, &answer.QuestionId)
		if err != nil {
			return nil, err
		}
		result = append(result, &answer)
	}

	return result, nil
}
