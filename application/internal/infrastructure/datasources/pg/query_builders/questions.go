package query_builders

import "github.com/Masterminds/squirrel"

const (
	questionsTableName     = "test_questions"
	questionsIdFieldName   = "id"
	questionsTextFieldName = "text"
)

func NewSelectQuestionByIdQuery(builder *squirrel.StatementBuilderType, id int) (string, []any, error) {
	return selectStarQuestion(builder).
		Where(squirrel.Eq{questionsIdFieldName: id}).
		ToSql()
}

func NewSelectAllQuestionsQuery(builder *squirrel.StatementBuilderType) (string, []any, error) {
	return selectStarQuestion(builder).
		ToSql()
}

func selectStarQuestion(builder *squirrel.StatementBuilderType) squirrel.SelectBuilder {
	return builder.
		Select(questionsIdFieldName, questionsTextFieldName).
		From(questionsTableName)
}
