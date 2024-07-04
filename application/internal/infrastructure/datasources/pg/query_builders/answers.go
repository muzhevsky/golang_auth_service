package query_builders

import (
	"github.com/Masterminds/squirrel"
)

func NewSelectAnswerValuesByAnswerIdQuery(builder *squirrel.StatementBuilderType, answerId int) (string, []any, error) {
	return builder.
		Select("id", "skill_id", "points").
		From("test_answers_values").
		Where(squirrel.Eq{"answer_id": answerId}).
		ToSql()
}

func NewSelectAnswersByQuestionIdQuery(builder *squirrel.StatementBuilderType, questionId int) (string, []any, error) {
	return builder.
		Select("id", "text", "question_id").
		From("test_answers").
		Where(squirrel.Eq{"question_id": questionId}).
		ToSql()
}

func NewSelectAnswerByIdQuery(builder *squirrel.StatementBuilderType, id int) (string, []any, error) {
	return builder.
		Select("text", "question_id").
		From("test_answers").
		Where(squirrel.Eq{"id": id}).
		ToSql()
}
