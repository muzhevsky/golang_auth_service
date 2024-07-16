package query_builders

import (
	"github.com/Masterminds/squirrel"
)

const (
	answerValuesTableName         = "test_answers_values"
	answerValuesIdFieldName       = "id"
	answerValuesAnswerIdFieldName = "answer_id"
	answerValuesSkillIdFieldName  = "skill_id"
	answerValuesPointsFieldName   = "points"
)

func NewSelectAnswerValuesByAnswerIdQuery(builder *squirrel.StatementBuilderType, answerId int) (string, []any, error) {
	return builder.
		Select(answerValuesIdFieldName, answerValuesAnswerIdFieldName, answerValuesSkillIdFieldName, answerValuesPointsFieldName).
		From(answerValuesTableName).
		Where(squirrel.Eq{answerValuesAnswerIdFieldName: answerId}).
		ToSql()
}
