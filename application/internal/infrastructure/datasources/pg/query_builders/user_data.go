package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/user_data"
)

func NewInsertUserDataQuery(builder *squirrel.StatementBuilderType, data *user_data.UserData) (string, []any, error) {
	return builder.
		Insert("user_data").
		Columns("nickname", "age", "gender", "xp", "account_id").
		Values(data.Nickname, data.Age, data.Gender, 0, data.AccountId).
		ToSql()
}

func NewSelectUserDataByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select("nickname", "age", "gender", "XP", "account_id").
		From("user_data").
		Where(squirrel.Eq{"account_id": accountId}).
		ToSql()
}

func NewUpdateUserDataByAccountIdQuery(builder *squirrel.StatementBuilderType, data *user_data.UserData) (string, []any, error) {
	return builder.Update("user_data").
		Set("nickname", data.Nickname).
		Set("xp", data.XP).
		Set("age", data.Age).
		Set("gender", data.Gender).
		Where(squirrel.Eq{"account_id": data.AccountId}).ToSql()
}