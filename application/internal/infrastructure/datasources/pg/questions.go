package pg

import (
	"context"
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities"
	"smartri_app/pkg/postgres"
)

type questionsDataSource struct {
	client *postgres.Client
}

func NewQuestionsDataSource(client *postgres.Client) *questionsDataSource {
	return &questionsDataSource{client: client}
}

func (q *questionsDataSource) SelectAll(context context.Context) ([]*entities.Question, error) {
	sql, args, err := q.client.Builder.
		Select("id", "text").
		From("test_questions").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := q.client.Pool.Query(context, sql, args...)

	if err != nil {
		return nil, err
	}

	result := make([]*entities.Question, 0)
	for rows.Next() {
		question := entities.Question{}
		err := rows.Scan(&question.Id, &question.Text)
		if err != nil {
			return nil, err
		}
		result = append(result, &question)
	}

	return result, nil
}

func (q *questionsDataSource) SelectById(context context.Context, id int) (*entities.Question, error) {
	sql, args, err := q.client.Builder.
		Select("id", "text").
		From("test_questions").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := q.client.Pool.QueryRow(context, sql, args...)

	var result entities.Question
	err = row.Scan(&result.Id, &result.Text)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
