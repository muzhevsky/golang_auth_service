package sessions

import (
	"authorization/internal/infrastructure/datasources"
	"context"
	"github.com/redis/go-redis/v9"
)

type deleteSessionByAccessTokenCommand struct {
	client *redis.Client
}

func NewDeleteSessionByAccessTokenCommand(client *redis.Client) datasources.IDeleteSessionByAccessTokenCommand {
	return &deleteSessionByAccessTokenCommand{client: client}
}

func (c *deleteSessionByAccessTokenCommand) Execute(context context.Context, accessToken string) error {
	key := getKey(accessToken)
	err := c.client.Del(context, key).Err()
	return err
}
