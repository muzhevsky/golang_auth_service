package query_builders

import (
	"authorization/internal/entities/session"
	sq "github.com/Masterminds/squirrel"
)

const (
	sessionsTableName             = "sessions"
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

func NewSelectSessionsByAccountIdQuery(builder *sq.StatementBuilderType, accountId int) (string, []any, error) {
	return selectStarSessions(builder).
		Where(sq.Eq{sessionsAccountIdFieldName: accountId}).
		ToSql()
}

func NewInsertSessionQuery(builder *sq.StatementBuilderType, session *session.Session) (string, []any, error) {
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
		ToSql()
}

func NewUpdateSessionQuery(builder *sq.StatementBuilderType, accessToken string, session *session.Session) (string, []any, error) {
	return builder.Update(sessionsTableName).
		Set(sessionsAccessTokenFieldName, session.AccessToken).
		Set(sessionsRefreshTokenFieldName, session.RefreshToken).
		Set(sessionsAccountIdFieldName, session.AccountId).
		Set(sessionsExpiresAtFieldName, session.ExpiresAt).
		Where(sq.Eq{sessionsAccessTokenFieldName: accessToken}).
		ToSql()
}

func selectStarSessions(builder *sq.StatementBuilderType) sq.SelectBuilder {
	return builder.
		Select(
			sessionsAccessTokenFieldName,
			sessionsRefreshTokenFieldName,
			sessionsAccountIdFieldName,
			sessionsExpiresAtFieldName).
		From(sessionsTableName)
}
