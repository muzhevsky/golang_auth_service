package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/test"
)

func NewInsertUserTestResultsQuery(builder *squirrel.StatementBuilderType, results *test.UserTestAnswers) (string, []any, error) {
	query := builder.Insert("user_answers").
		Columns("account_id", "question_id", "answer_id")
	for i := range results.Answers {
		query = query.Values(results.AccountId, results.Answers[i].QuestionId, results.Answers[i].AnswerId)
	}

	return query.ToSql()
}
