package pg

import (
	"context"
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities"
	"smartri_app/pkg/postgres"
)

type answersDataSource struct {
	client *postgres.Client
}

func NewAnswersDataSource(client *postgres.Client) *answersDataSource {
	return &answersDataSource{client: client}
}

func (q *answersDataSource) SelectByQuestionId(context context.Context, questionId int) ([]*entities.Answer, error) {
	sql, args, err := q.client.Builder.
		Select("id", "text", "question_id").
		From("test_answers").
		Where(squirrel.Eq{"question_id": questionId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := q.client.Pool.Query(context, sql, args...)

	if err != nil {
		return nil, err
	}

	result := make([]*entities.Answer, 0)
	for rows.Next() {
		answer := entities.Answer{}
		err := rows.Scan(&answer.Id, &answer.Text, &answer.QuestionId)
		if err != nil {
			return nil, err
		}
		result = append(result, &answer)
	}

	return result, nil
}

func (q *answersDataSource) SelectById(context context.Context, id int) (*entities.Answer, error) {
	sql, args, err := q.client.Builder.
		Select("id", "text", "question_id").
		From("test_answers").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := q.client.Pool.QueryRow(context, sql, args...)

	result := entities.Answer{}
	err = row.Scan(&result.Id, &result.Text, &result.QuestionId)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
