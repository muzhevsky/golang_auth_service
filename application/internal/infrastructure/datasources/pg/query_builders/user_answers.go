package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/test_entities"
)

const (
	userAnswersTableName           = "user_answers"
	userAnswersIdFieldName         = "id"
	userAnswersAccountIdFieldName  = "account_id"
	userAnswersQuestionIdFieldName = "question_id"
	userAnswersAnswerIdFieldName   = "answer_id"
)

func NewInsertUserTestResultsQuery(builder *squirrel.StatementBuilderType, results *test_entities.UserTestAnswers) (string, []any, error) {
	query := builder.
		Insert(userAnswersTableName).
		Columns(userAnswersAccountIdFieldName, userAnswersQuestionIdFieldName, userAnswersAnswerIdFieldName)
	for i := range results.Answers {
		query = query.Values(results.AccountId, results.Answers[i].QuestionId, results.Answers[i].AnswerId)
	}

	return query.ToSql()
}

func NewSelectUserAnswersByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return selectStarUserAnswers(builder).
		Where(squirrel.Eq{userAnswersIdFieldName: accountId}).
		ToSql()
}

func NewSelectUserAnswerByAccountIdQueryLimit1(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return selectStarUserAnswers(builder).
		Where(squirrel.Eq{"account_id": accountId}).
		Limit(1).
		ToSql()
}

func selectStarUserAnswers(builder *squirrel.StatementBuilderType) squirrel.SelectBuilder {
	return builder.
		Select(
			userAnswersIdFieldName,
			userAnswersQuestionIdFieldName,
			userAnswersAnswerIdFieldName,
			userAnswersAccountIdFieldName).
		From(userAnswersTableName)
}
