package verification

import (
	"authorization/internal/entities/verification"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/redis/commands"
	"context"
	"github.com/redis/go-redis/v9"
)

type selectVerificationByAccountIdRedisCommand struct {
	client *redis.Client
}

func NewSelectVerificationByAccountIdRedisCommand(client *redis.Client) datasources.ISelectVerificationsByAccountIdCommand {
	return &selectVerificationByAccountIdRedisCommand{client: client}
}

func (c *selectVerificationByAccountIdRedisCommand) Execute(context context.Context, accountId int) ([]*verification.Verification, error) {
	key := getKey(accountId)
	verificationsPtr, err := commands.GetValueOrNil[[]*verification.Verification](context, c.client, key)
	if err != nil {
		return nil, err
	}

	if verificationsPtr == nil {
		return make([]*verification.Verification, 0), nil
	}

	verifications := *verificationsPtr

	return verifications, nil
}
