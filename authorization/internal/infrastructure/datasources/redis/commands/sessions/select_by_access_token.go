package sessions

import (
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/redis/commands"
	"context"
	"github.com/redis/go-redis/v9"
)

type selectSessionByAccessTokenRedisCommand struct {
	client *redis.Client
}

func NewSelectSessionByAccessTokenRedisCommand(client *redis.Client) datasources.ISelectSessionByAccessTokenCommand {
	return &selectSessionByAccessTokenRedisCommand{client: client}
}

func (c *selectSessionByAccessTokenRedisCommand) Execute(context context.Context, accessToken string) (*session.Session, error) {
	key := getKey(accessToken)
	value, err := commands.GetValueOrNil[session.Session](context, c.client, key)
	if err != nil {
		return nil, err
	}

	return value, nil
}
