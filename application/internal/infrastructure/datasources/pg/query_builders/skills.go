package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities"
)

func NewSelectAllSkillsQuery(builder *squirrel.StatementBuilderType) (string, []any, error) {
	return builder.
		Select("id", "title").
		From("skills").
		ToSql()
}

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

func NewSelectSkillChangesByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select("id", "skill_id", "date", "action_id", "points").
		From("skill_changes").
		Where(squirrel.Eq{"account_id": accountId}).
		ToSql()
}

func NewSelectSkillChangesByAccountIdAndActionIdQuery(builder *squirrel.StatementBuilderType, accountId int, actionId int) (string, []any, error) {
	return builder.
		Select("id", "skill_id", "date", "points").
		From("skill_changes").
		Where(squirrel.Eq{"account_id": accountId, "action_id": actionId}).
		ToSql()
}

func NewInsertSkillChangesQuery(builder *squirrel.StatementBuilderType, change *entities.SkillChange) (string, []any, error) {
	return builder.
		Insert("skill_changes").
		Columns("account_id", "skill_id", "date", "action_id", "points").
		Values(change.AccountId, change.SkillId, change.Date, change.ActionId, change.Points).
		ToSql()
}

func NewSelectSkillsByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select("skill_id", "xp").
		From("user_skills").
		Where(squirrel.Eq{"account_id": accountId}).
		ToSql()
}

func NewUpdateUserSkillsByAccountIdQuery(builder *squirrel.StatementBuilderType, skill *entities.UserSkills) (string, []any, error) {
	return builder.
		Update("user_skills").
		Set("xp", skill.Xp).
		Where(squirrel.Eq{"account_id": skill.AccountId, "skill_id": skill.SkillId}).
		ToSql()
}
