package sessions

import (
	"authorization/internal/entities/session_entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/redis/commands"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type updateSessionByAccessTokenRedisCommand struct {
	client *redis.Client
}

func NewUpdateSessionByAccessTokenRedisCommand(client *redis.Client) datasources.IUpdateSessionByAccessTokenCommand {
	return &updateSessionByAccessTokenRedisCommand{client: client}
}

func (c updateSessionByAccessTokenRedisCommand) Execute(context context.Context, accessToken string, newSession *session_entities.Session) error {
	oldKey := getKey(accessToken)
	s, err := commands.GetValueOrNil[session_entities.Session](context, c.client, oldKey)
	if err != nil {
		return err
	}
	if s != nil {
		err = c.client.Del(context, oldKey).Err()
		if err != nil {
			return err
		}
	}

	newKey := getKey(newSession.AccessToken)

	err = commands.SetValue(context, c.client, newKey, newSession, newSession.ExpiresAt.Sub(time.Now()))
	if err != nil {
		return err
	}

	return nil
}
