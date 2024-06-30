package query_builders

import "github.com/Masterminds/squirrel"

func NewSelectQuestionByIdQuery(builder *squirrel.StatementBuilderType, id int) (string, []any, error) {
	return builder.
		Select("id", "text").
		From("test_questions").
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func NewSelectAllQuestionsQuery(builder *squirrel.StatementBuilderType) (string, []any, error) {
	return builder.
		Select("id", "text").
		From("test_questions").
		ToSql()
}
