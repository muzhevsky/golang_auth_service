package query_builders

import "github.com/Masterminds/squirrel"

const (
	skillNormalizationsTableName        = "skill_normalizations"
	skillNormalizationsSkillIdFieldName = "skill_Id"
	skillNormalizationsMinimumFieldName = "minimum"
	skillNormalizationsMaximumFieldName = "maximum"
)

func NewSelectAllSkillsNormalizationsQuery(builder *squirrel.StatementBuilderType) (string, []any, error) {
	return selectStarSkillNormalization(builder).
		From("skill_normalizations").
		ToSql()
}

func NewSelectSkillNormalizationsBySkillIdQuery(builder *squirrel.StatementBuilderType, skillId int) (string, []any, error) {
	return selectStarSkillNormalization(builder).
		From("skill_normalizations").
		Where(squirrel.Eq{"skill_id": skillId}).
		ToSql()
}

func selectStarSkillNormalization(builder *squirrel.StatementBuilderType) squirrel.SelectBuilder {
	return builder.
		Select(
			skillNormalizationsSkillIdFieldName,
			skillNormalizationsMinimumFieldName,
			skillNormalizationsMaximumFieldName).
		From(skillNormalizationsTableName)
}
