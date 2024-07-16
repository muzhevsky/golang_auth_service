package sessions

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/redis/commands"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type insertSessionRedisCommand struct {
	client *redis.Client
}

func NewInsertSessionRedisCommand(client *redis.Client) datasources.IInsertSessionCommand {
	return &insertSessionRedisCommand{client: client}
}

func (c *insertSessionRedisCommand) Execute(ctx context.Context, s *session_entities.Session) error {
	key := getKey(s.AccessToken)
	sessionPtr, err := commands.GetValueOrNil[session_entities.Session](ctx, c.client, key)
	if err != nil {
		return err
	}
	if sessionPtr != nil {
		return err // TODO
	}

	err = commands.SetValue(ctx, c.client, key, s, s.ExpiresAt.Sub(time.Now()))
	return err
}
