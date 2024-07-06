package query_builders

import (
	"authorization/internal/entities"
	sq "github.com/Masterminds/squirrel"
)

const (
	sessionsTableName             = "sessions"
	sessionsIdFieldName           = "id"
	sessionsAccessTokenFieldName  = "access_token"
	sessionsRefreshTokenFieldName = "refresh_token"
	sessionsAccountIdFieldName    = "account_id"
	sessionsExpiresAtFieldName    = "expires_at"
)

func NewSelectSessionByAccessTokenQuery(builder *sq.StatementBuilderType, accessToken string) (string, []any, error) {
	return selectStarSessions(builder).
		Where(sq.Eq{sessionsAccessTokenFieldName: accessToken}).
		ToSql()
}

func NewSelectSessionByIdQuery(builder *sq.StatementBuilderType, id int) (string, []any, error) {
	return selectStarSessions(builder).
		Where(sq.Eq{sessionsIdFieldName: id}).
		ToSql()
}

func NewSelectSessionsByAccountIdQuery(builder *sq.StatementBuilderType, accountId int) (string, []any, error) {
	return selectStarSessions(builder).
		Where(sq.Eq{sessionsAccountIdFieldName: accountId}).
		ToSql()
}

func NewInsertSessionQuery(builder *sq.StatementBuilderType, session *entities.Session) (string, []any, error) {
	return builder.
		Insert(sessionsTableName).
		Columns(
			sessionsAccessTokenFieldName,
			sessionsRefreshTokenFieldName,
			sessionsAccountIdFieldName,
			sessionsExpiresAtFieldName).
		Values(
			session.AccessToken,
			session.RefreshToken,
			session.AccountId,
			session.ExpiresAt).
		Suffix("returning " + sessionsIdFieldName).
		ToSql()
}

func NewUpdateSessionQuery(builder *sq.StatementBuilderType, session *entities.Session) (string, []any, error) {
	return builder.Update(sessionsTableName).
		Set(sessionsAccessTokenFieldName, session.AccessToken).
		Set(sessionsRefreshTokenFieldName, session.RefreshToken).
		Set(sessionsAccountIdFieldName, session.AccountId).
		Set(sessionsExpiresAtFieldName, session.ExpiresAt).
		Where(sq.Eq{sessionsIdFieldName: session.Id}).
		ToSql()
}

func selectStarSessions(builder *sq.StatementBuilderType) sq.SelectBuilder {
	return builder.
		Select(sessionsIdFieldName,
			sessionsAccessTokenFieldName,
			sessionsRefreshTokenFieldName,
			sessionsAccountIdFieldName,
			sessionsExpiresAtFieldName).
		From(sessionsTableName)
}
