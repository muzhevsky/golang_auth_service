package sessions

import (
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources/redis/commands"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type replaceSessionByAccessTokenRedisCommand struct {
	client *redis.Client
}

func NewReplaceSessionByAccessTokenRedisCommand(client *redis.Client) *replaceSessionByAccessTokenRedisCommand {
	return &replaceSessionByAccessTokenRedisCommand{client: client}
}

func (c replaceSessionByAccessTokenRedisCommand) Execute(context context.Context, accessToken string, newSession *session.Session) error {
	oldKey := getKey(accessToken)
	session, err := commands.GetValueOrNil[session.Session](context, c.client, oldKey)
	if err != nil {
		return err
	}
	if session != nil {
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
