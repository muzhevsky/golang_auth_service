package query_builders

import (
	"github.com/Masterminds/squirrel"
)

func NewSelectAllSkillsQuery(builder *squirrel.StatementBuilderType) (string, []any, error) {
	return builder.
		Select("id", "title").
		From("skills").
		ToSql()
}
