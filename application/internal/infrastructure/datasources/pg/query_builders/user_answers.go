package query_builders

import "github.com/Masterminds/squirrel"

func NewSelectUserAnswersByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select("question_id", "answer_id").
		From("user_answers").
		Where(squirrel.Eq{"id": accountId}).
		ToSql()
}

func NewSelectUserAnswerByAccountIdQueryLimit1(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select("question_id", "answer_id").
		From("user_answers").
		Where(squirrel.Eq{"account_id": accountId}).
		Limit(1).
		ToSql()
}
