package factories

import (
	"authorization/internal"
	"authorization/internal/infrastructure/datasources/pg/commands/accounts"
	"authorization/internal/infrastructure/datasources/pg/commands/devices"
	sessions2 "authorization/internal/infrastructure/datasources/redis/commands/sessions"
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

func CreateSessionRepo(redis *redis.Client) internal.ISessionRepository {
	selectSessionByAccessTokenCommand := sessions2.NewSelectSessionByAccessTokenRedisCommand(redis)
	updateSessionByAccessTokenCommand := sessions2.NewUpdateSessionByAccessTokenRedisCommand(redis)
	insertSessionCommand := sessions2.NewInsertSessionRedisCommand(redis)
	deleteSessionByAccessTokenCommand := sessions2.NewDeleteSessionByAccessTokenCommand(redis)

	return repositories.NewSessionRepository(
		selectSessionByAccessTokenCommand,
		insertSessionCommand,
		updateSessionByAccessTokenCommand,
		deleteSessionByAccessTokenCommand)
}

func CreateDeviceRepo(postgres *postgres.Client) internal.IDeviceRepository {
	selectDeviceByIdCommand := devices.NewSelectDeviceByIdPGCommand(postgres)
	selectDevicesByAccountIdCommand := devices.NewSelectDevicesByAccountIdPGCommand(postgres)
	deleteDeviceByIdCommand := devices.NewDeleteDeviceByIdPGCommand(postgres)

	selectDeviceByAccessTokenCommand := devices.NewSelectDeviceByAccessTokenPGCommand(postgres)
	updateDeviceByAccessTokenCommand := devices.NewUpdateDeviceByAccessTokenPGCommand(postgres)
	insertDeviceCommand := devices.NewInsertDevicePGCommand(postgres)

	return repositories.NewDeviceRepo(
		insertDeviceCommand,
		selectDevicesByAccountIdCommand,
		selectDeviceByIdCommand,
		selectDeviceByAccessTokenCommand,
		updateDeviceByAccessTokenCommand,
		deleteDeviceByIdCommand)
}
