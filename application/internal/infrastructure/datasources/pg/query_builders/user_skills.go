package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/skills_entities"
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

func NewUpdateUserSkillsByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int, skill *skills_entities.UserSkill) (string, []any, error) {
	return builder.
		Update(userSkillsTableName).
		Set(userSkillXpFieldName, skill.Xp).
		Where(squirrel.Eq{"account_id": accountId, userSkillSkillIdFieldName: skill.SkillId}).
		ToSql()
}

func NewInsertUserSkillsQuery(builder *squirrel.StatementBuilderType, accountId int, userSkills *skills_entities.UserSkill) (string, []any, error) {
	return builder.Insert("user_skills").
		Columns("account_id", "skill_id", "xp").
		Values(accountId, userSkills.SkillId, userSkills.Xp).
		ToSql()
}

func NewUpdateUserSkillsQuery(builder *squirrel.StatementBuilderType, accountId int, skills *skills_entities.UserSkill) (string, []any, error) {
	return builder.Update("user_skills").
		Set("xp", skills.Xp).
		Where(squirrel.Eq{"account_id": accountId, "skill_id": skills.SkillId}).
		ToSql()
}
