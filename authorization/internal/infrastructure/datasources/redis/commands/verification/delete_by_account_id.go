package verification

import (
	"authorization/internal/infrastructure/datasources"
	"context"
	"github.com/redis/go-redis/v9"
)

type deleteVerificationByAccountIdRedisCommand struct {
	client *redis.Client
}

func NewDeleteVerificationByAccountIdRedisCommand(client *redis.Client) datasources.IDeleteVerificationsByAccountIdCommand {
	return &deleteVerificationByAccountIdRedisCommand{client: client}
}

func (c *deleteVerificationByAccountIdRedisCommand) Execute(context context.Context, accountId int) error {
	key := getKey(accountId)
	err := c.client.Del(context, key).Err()
	return err
}
