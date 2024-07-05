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

func NewSelectAccountByIdQuery(builder *sq.StatementBuilderType, id int) (string, []any, error) {
	return builder.Select(
		accountIdFieldName,
		accountLoginFieldName,
		accountPasswordFieldName,
		accountNicknameFieldName,
		accountEmailFieldName,
		accountRegistrationDateFieldName,
		accountIsVerifiedFieldName).
		From(accountTableName).
		Where(sq.Eq{accountIdFieldName: id}).
		ToSql()
}

func NewSelectAccountByLoginQuery(builder *sq.StatementBuilderType, login string) (string, []any, error) {
	return builder.Select(
		accountIdFieldName,
		accountLoginFieldName,
		accountPasswordFieldName,
		accountNicknameFieldName,
		accountEmailFieldName,
		accountRegistrationDateFieldName,
		accountIsVerifiedFieldName).
		From(accountTableName).
		Where(sq.Eq{accountLoginFieldName: login}).
		ToSql()
}

func NewSelectAccountByEmailQuery(builder *sq.StatementBuilderType, email string) (string, []any, error) {
	return builder.
		Select(
			accountIdFieldName,
			accountLoginFieldName,
			accountPasswordFieldName,
			accountNicknameFieldName,
			accountEmailFieldName,
			accountRegistrationDateFieldName,
			accountIsVerifiedFieldName).
		From(accountTableName).
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
