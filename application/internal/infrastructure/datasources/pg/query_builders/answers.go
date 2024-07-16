package query_builders

import "github.com/Masterminds/squirrel"

const (
	answersTableName           = "test_answers"
	answersIdFieldName         = "id"
	answersTextFieldName       = "text"
	answersQuestionIdFieldName = "question_id"
)

func NewSelectAnswersByQuestionIdQuery(builder *squirrel.StatementBuilderType, questionId int) (string, []any, error) {
	return selectStartAnswer(builder).
		Where(squirrel.Eq{answersQuestionIdFieldName: questionId}).
		ToSql()
}

func NewSelectAnswerByIdQuery(builder *squirrel.StatementBuilderType, id int) (string, []any, error) {
	return selectStartAnswer(builder).
		Where(squirrel.Eq{answersIdFieldName: id}).
		ToSql()
}

func selectStartAnswer(builder *squirrel.StatementBuilderType) squirrel.SelectBuilder {
	return builder.
		Select(answersIdFieldName, answersTextFieldName, answersQuestionIdFieldName).
		From(answersTableName)
}
