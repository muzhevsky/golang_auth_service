package answers

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAnswersByQuestionIdCommand struct {
	client *postgres.Client
}

func NewSelectAnswersByQuestionIdCommand(client *postgres.Client) *selectAnswersByQuestionIdCommand {
	return &selectAnswersByQuestionIdCommand{client: client}
}

func (q *selectAnswersByQuestionIdCommand) Execute(context context.Context, questionId int) ([]*entities.Answer, error) {
	sql, args, err := query_builders.NewSelectAnswersByQuestionIdQuery(&q.client.Builder, questionId)
	if err != nil {
		return nil, err
	}

	rows, err := q.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*entities.Answer, 0)
	for rows.Next() {
		answer := entities.Answer{}
		err = rows.Scan(&answer.Id, &answer.Text, &answer.QuestionId)
		if err != nil {
			return nil, err
		}
		result = append(result, &answer)
	}

	return result, nil
}
