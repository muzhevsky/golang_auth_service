package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/skills"
)

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

func NewInsertSkillChangesQuery(builder *squirrel.StatementBuilderType, change *skills.SkillChange) (string, []any, error) {
	return builder.
		Insert("skill_changes").
		Columns("account_id", "skill_id", "date", "action_id", "points").
		Values(change.AccountId, change.SkillId, change.Date, change.ActionId, change.Points).
		ToSql()
}
