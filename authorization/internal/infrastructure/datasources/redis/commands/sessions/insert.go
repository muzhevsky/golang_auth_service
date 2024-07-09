package sessions

import (
	"authorization/internal/entities/session"
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

func (c *insertSessionRedisCommand) Execute(ctx context.Context, session *session.Session) error {
	key := getKey(session.AccessToken)
	sessionPtr, err := commands.GetValueOrNil[session.Session](ctx, c.client, key)
	if err != nil {
		return err
	}
	if sessionPtr != nil {
		return err // TODO
	}

	err = commands.SetValue(ctx, c.client, key, session, session.ExpiresAt.Sub(time.Now()))
	return err
}
