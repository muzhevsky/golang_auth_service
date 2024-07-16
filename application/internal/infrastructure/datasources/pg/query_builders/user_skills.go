package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/skills_entities"
)

const (
	userSkillsTableName         = "user_skills"
	userSkillXpFieldName        = "xp"
	userSkillAccountIdFieldName = "account_id"
	userSkillSkillIdFieldName   = "skill_id"
)

func NewSelectUserSkillsByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select(userSkillSkillIdFieldName, userSkillXpFieldName).
		From(userSkillsTableName).
		Where(squirrel.Eq{userSkillAccountIdFieldName: accountId}).
		ToSql()
}

func NewUpdateUserSkillsByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int, skill *skills_entities.UserSkill) (string, []any, error) {
	return builder.
		Update(userSkillsTableName).
		Set(userSkillXpFieldName, skill.Xp).
		Where(squirrel.Eq{userSkillAccountIdFieldName: accountId, userSkillSkillIdFieldName: skill.SkillId}).
		ToSql()
}

func NewInsertUserSkillsQuery(builder *squirrel.StatementBuilderType, accountId int, userSkills *skills_entities.UserSkill) (string, []any, error) {
	return builder.Insert(userSkillsTableName).
		Columns(userSkillAccountIdFieldName, userSkillSkillIdFieldName, userSkillXpFieldName).
		Values(accountId, userSkills.SkillId, userSkills.Xp).
		ToSql()
}

func NewUpdateUserSkillsQuery(builder *squirrel.StatementBuilderType, accountId int, skills *skills_entities.UserSkill) (string, []any, error) {
	return builder.Update(userSkillsTableName).
		Set(userSkillXpFieldName, skills.Xp).
		Where(squirrel.Eq{userSkillAccountIdFieldName: accountId, userSkillSkillIdFieldName: skills.SkillId}).
		ToSql()
}
