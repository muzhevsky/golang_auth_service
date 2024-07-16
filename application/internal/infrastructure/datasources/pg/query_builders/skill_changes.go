package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/skills_entities"
)

const (
	skillChangesTableName         = "skill_changes"
	skillChangesAccountIdField    = "account_id"
	skillChangesIdFieldName       = "id"
	skillChangesSkillIdFieldName  = "skill_id"
	skillChangesDateFieldName     = "date"
	skillChangesActionIdFieldName = "action_id"
	skillChangesPointsFieldName   = "points"
)

func NewSelectSkillChangesByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return selectStarSkillChanges(builder).
		Where(squirrel.Eq{skillChangesAccountIdField: accountId}).
		ToSql()
}

func NewSelectSkillChangesByAccountIdAndActionIdQuery(builder *squirrel.StatementBuilderType, accountId int, actionId int) (string, []any, error) {
	return selectStarSkillChanges(builder).
		Where(squirrel.Eq{skillChangesAccountIdField: accountId, skillChangesActionIdFieldName: actionId}).
		ToSql()
}

func NewInsertSkillChangesQuery(builder *squirrel.StatementBuilderType, change *skills_entities.SkillChange) (string, []any, error) {
	return builder.
		Insert("skill_changes").
		Columns("account_id", skillChangesSkillIdFieldName, skillChangesDateFieldName, skillChangesActionIdFieldName, "points").
		Values(change.AccountId, change.SkillId, change.Date, change.ActionId, change.Points).
		ToSql()
}

func selectStarSkillChanges(builder *squirrel.StatementBuilderType) squirrel.SelectBuilder {
	return builder.Select(
		skillChangesIdFieldName,
		skillChangesAccountIdField,
		skillChangesSkillIdFieldName,
		skillChangesDateFieldName,
		skillChangesActionIdFieldName,
		skillChangesPointsFieldName).
		From(skillChangesTableName)
}
