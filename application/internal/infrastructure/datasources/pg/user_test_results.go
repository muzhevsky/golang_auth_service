package pg

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/pkg/postgres"
)

type userTestResults struct {
	client *postgres.Client
}

func NewUserTestResultsSource(client *postgres.Client) *userTestResults {
	return &userTestResults{client: client}
}

func (u *userTestResults) Insert(context context.Context, results *entities.UserTestAnswers) error {
	query := u.client.Builder.Insert("user_answers").
		Columns("account_id", "question_id", "answer_id")
	for i := range results.Answers {
		query = query.Values(results.AccountId, results.Answers[i].QuestionId, results.Answers[i].AnswerId)
	}

	sql, i, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = u.client.Pool.Exec(context, sql, i...)
	if err != nil {
		return err
	}

	return nil
}
