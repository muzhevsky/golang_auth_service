package query_builders

import (
	"authorization/internal/entities/account"
	sq "github.com/Masterminds/squirrel"
)

const (
	accountTableName                 = "accounts"
	accountIdFieldName               = "id"
	accountLoginFieldName            = "login"
	accountEmailFieldName            = "email"
	accountPasswordFieldName         = "password"
	accountNicknameFieldName         = "nickname"
	accountIsVerifiedFieldName       = "is_verified"
	accountRegistrationDateFieldName = "registration_date"
)

// NewSelectAccountByIdQuery - creates sql string with args and error using squirrel StatementBuilder
// order of selected fields:
//   - id
//   - login
//   - password
//   - nickname
//   - email
//   - registration_date
//   - is_verified
func NewSelectAccountByIdQuery(builder *sq.StatementBuilderType, id int) (string, []any, error) {
	return selectStarAccounts(builder).
		Where(sq.Eq{accountIdFieldName: id}).
		ToSql()
}

// NewSelectAccountByLoginQuery - creates sql string with args and error using squirrel StatementBuilder
// order of selected fields:
//   - id
//   - login
//   - password
//   - nickname
//   - email
//   - registration_date
//   - is_verified
func NewSelectAccountByLoginQuery(builder *sq.StatementBuilderType, login string) (string, []any, error) {
	return selectStarAccounts(builder).
		Where(sq.Eq{accountLoginFieldName: login}).
		ToSql()
}

// NewSelectAccountByEmailQuery - creates sql string with args and error using squirrel StatementBuilder
// order of selected fields:
//   - id
//   - login
//   - password
//   - nickname
//   - email
//   - registration_date
//   - is_verified
func NewSelectAccountByEmailQuery(builder *sq.StatementBuilderType, email string) (string, []any, error) {
	return selectStarAccounts(builder).
		Where(sq.Eq{accountEmailFieldName: email}).
		ToSql()
}

func NewInsertAccountQuery(builder *sq.StatementBuilderType, account *account.Account) (string, []any, error) {
	return builder.Insert(accountTableName).
		Columns(
			accountLoginFieldName,
			accountEmailFieldName,
			accountNicknameFieldName,
			accountPasswordFieldName,
			accountRegistrationDateFieldName,
			accountIsVerifiedFieldName).
		Values(account.Login, account.Email, account.Nickname, account.Password, account.RegistrationDate, account.IsVerified).
		Suffix("RETURNING " + accountIdFieldName).
		ToSql()
}

func NewUpdateAccountByIdQuery(builder *sq.StatementBuilderType, id int, newAccount *account.Account) (string, []any, error) {
	return builder.Update(accountTableName).
		Set(accountLoginFieldName, newAccount.Login).
		Set(accountEmailFieldName, newAccount.Email).
		Set(accountNicknameFieldName, newAccount.Nickname).
		Set(accountPasswordFieldName, newAccount.Password).
		Set(accountRegistrationDateFieldName, newAccount.RegistrationDate).
		Set(accountIsVerifiedFieldName, newAccount.IsVerified).
		Where(sq.Eq{accountIdFieldName: id}).
		ToSql()
}

func selectStarAccounts(builder *sq.StatementBuilderType) sq.SelectBuilder {
	return builder.
		Select(
			accountIdFieldName,
			accountLoginFieldName,
			accountPasswordFieldName,
			accountNicknameFieldName,
			accountEmailFieldName,
			accountRegistrationDateFieldName,
			accountIsVerifiedFieldName).
		From(accountTableName)
}
