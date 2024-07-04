package query_builders

import "github.com/Masterminds/squirrel"

func NewSelectAllSkillsNormalizationsQuery(builder *squirrel.StatementBuilderType) (string, []any, error) {
	return builder.
		Select("skill_id", "minimum", "maximum").
		From("skill_normalizations").
		ToSql()
}

func NewSelectSkillNormalizationsBySkillIdQuery(builder *squirrel.StatementBuilderType, skillId int) (string, []any, error) {
	return builder.
		Select("skill_id", "minimum", "maximum").
		From("skill_normalizations").
		Where(squirrel.Eq{"skill_id": skillId}).
		ToSql()
}
