package factories

import (
	"authorization/internal"
	"authorization/internal/infrastructure/datasources/pg/commands/accounts"
	"authorization/internal/infrastructure/datasources/pg/commands/sessions"
	"authorization/internal/infrastructure/datasources/redis/commands/verification"
	"authorization/internal/repositories"
	"authorization/pkg/postgres"
	"github.com/redis/go-redis/v9"
)

func CreateRedisVerificationRepo(client *redis.Client) internal.IVerificationRepository {
	createVerificationCommand := verification.NewCreateVerificationRedisCommand(client)
	selectVerificationsByAccountIdCommand := verification.NewSelectVerificationByAccountIdRedisCommand(client)
	deleteVerificationsByAccountIdCommand := verification.NewDeleteVerificationByAccountIdRedisCommand(client)
	return repositories.NewVerificationRepo(
		createVerificationCommand,
		selectVerificationsByAccountIdCommand,
		deleteVerificationsByAccountIdCommand)
}

func CreatePGAccountRepo(client *postgres.Client) internal.IAccountRepository {
	selectAccountByIdCommand := accounts.NewSelectAccountByIdPGCommand(client)
	selectAccountByEmailCommand := accounts.NewSelectAccountByEmailPGCommand(client)
	selectAccountByLoginCommand := accounts.NewSelectAccountByLoginPGCommand(client)
	updateAccountByIdCommand := accounts.NewUpdateAccountByIdPGCommand(client)
	insertAccountCommand := accounts.NewInsertAccountPGCommand(client)

	return repositories.NewAccountRepository(
		selectAccountByIdCommand,
		selectAccountByEmailCommand,
		selectAccountByLoginCommand,
		updateAccountByIdCommand,
		insertAccountCommand)
}

func CreatePGSessionRepo(client *postgres.Client) internal.ISessionRepository {
	selectSessionByIdCommand := sessions.NewSelectSessionByIdPGCommand(client)
	selectSessionByAccessTokenCommand := sessions.NewSelectSessionByAccessTokenPGCommand(client)
	selectSessionsByAccountIdCommand := sessions.NewSelectSessionByAccountIdPGCommand(client)
	insertSessionCommand := sessions.NewInsertSessionPGCommand(client)
	updateSessionByIdCommand := sessions.NewUpdateSessionByIdPGCommand(client)

	return repositories.NewSessionRepository(
		selectSessionByIdCommand,
		selectSessionByAccessTokenCommand,
		selectSessionsByAccountIdCommand,
		insertSessionCommand,
		updateSessionByIdCommand)
}
