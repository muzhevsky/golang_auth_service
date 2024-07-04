package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/user_data"
)

const userSkillsTableName = "user_skills"
const userSkillXpFieldName = "xp"
const userSkillSkillIdFieldName = "skill_id"

func NewSelectUserSkillsByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select(userSkillSkillIdFieldName, userSkillXpFieldName).
		From(userSkillsTableName).
		Where(squirrel.Eq{"account_id": accountId}).
		ToSql()
}

func NewUpdateUserSkillsByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int, skill *user_data.UserSkill) (string, []any, error) {
	return builder.
		Update(userSkillsTableName).
		Set(userSkillXpFieldName, skill.Xp).
		Where(squirrel.Eq{"account_id": accountId, userSkillSkillIdFieldName: skill.SkillId}).
		ToSql()
}
