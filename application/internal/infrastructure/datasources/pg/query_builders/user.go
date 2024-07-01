package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities"
)

func NewInsertUserDataQuery(builder *squirrel.StatementBuilderType, data *entities.UserData) (string, []any, error) {
	return builder.
		Insert("user_data").
		Columns("age", "gender", "xp", "account_id").
		Values(data.Age, data.Gender, 0, data.AccountId).
		ToSql()
}

func NewSelectUserDataByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select("age", "gender", "XP", "account_id").
		From("user_data").
		Where(squirrel.Eq{"account_id": accountId}).
		ToSql()
}

func NewUpdateUserDataByAccountIdQuery(builder *squirrel.StatementBuilderType, data *entities.UserData) (string, []any, error) {
	return builder.Update("user_data").
		Set("xp", data.XP).
		Set("age", data.Age).
		Set("gender", data.Gender).
		Where(squirrel.Eq{"account_id": data.AccountId}).ToSql()
}

func NewInsertUserSkillsQuery(builder *squirrel.StatementBuilderType, userSkills *entities.UserSkills) (string, []any, error) {
	return builder.Insert("user_skills").
		Columns("account_id", "skill_id", "xp").
		Values(userSkills.AccountId, userSkills.SkillId, userSkills.Xp).
		ToSql()
}

func NewUpdateUserSkillsQuery(builder *squirrel.StatementBuilderType, skills *entities.UserSkills) (string, []any, error) {
	return builder.Update("user_skills").
		Set("xp", skills.Xp).
		Where(squirrel.Eq{"account_id": skills.AccountId, "skill_id": skills.SkillId}).
		ToSql()
}
