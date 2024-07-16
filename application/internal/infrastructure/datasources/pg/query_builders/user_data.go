package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/user_data_entities"
)

const (
	userDataTableName          = "user_data"
	userDataNicknameFieldName  = "nickname"
	userDataAgeFieldName       = "age"
	userDataGenderFieldName    = "gender"
	userDataXPFieldName        = "xp"
	userDataAccountIdFieldName = "account_id"
)

func NewInsertUserDataQuery(builder *squirrel.StatementBuilderType, data *user_data_entities.UserData) (string, []any, error) {
	return builder.
		Insert(userDataTableName).
		Columns(userDataNicknameFieldName, userDataAgeFieldName, userDataGenderFieldName, userDataXPFieldName, userDataAccountIdFieldName).
		Values(data.Nickname, data.Age, data.Gender, 0, data.AccountId).
		ToSql()
}

func NewSelectUserDataByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select(userDataNicknameFieldName, userDataAgeFieldName, userDataGenderFieldName, userDataXPFieldName, userDataAccountIdFieldName).
		From(userDataTableName).
		Where(squirrel.Eq{userDataAccountIdFieldName: accountId}).
		ToSql()
}

func NewUpdateUserDataByAccountIdQuery(builder *squirrel.StatementBuilderType, data *user_data_entities.UserData) (string, []any, error) {
	return builder.Update(userDataTableName).
		Set(userDataNicknameFieldName, data.Nickname).
		Set(userDataAgeFieldName, data.Age).
		Set(userDataGenderFieldName, data.Gender).
		Set(userDataXPFieldName, data.XP).
		Where(squirrel.Eq{userDataAccountIdFieldName: data.AccountId}).ToSql()
}
