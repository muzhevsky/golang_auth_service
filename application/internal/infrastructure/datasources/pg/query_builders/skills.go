package query_builders

import (
	"github.com/Masterminds/squirrel"
)

const (
	skillsTableName      = "skills"
	skillsIdFieldName    = "id"
	skillsTitleFieldName = "title"
)

func NewSelectAllSkillsQuery(builder *squirrel.StatementBuilderType) (string, []any, error) {
	return builder.
		Select(skillsIdFieldName, skillsTitleFieldName).
		From(skillsTableName).
		ToSql()
}
