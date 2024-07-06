package verification

import (
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"authorization/internal/infrastructure/datasources/redis/commands"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type createVerificationRedisCommand struct {
	client *redis.Client
}

func NewCreateVerificationRedisCommand(client *redis.Client) datasources.ICreateVerificationCommand {
	return &createVerificationRedisCommand{client: client}
}

func (c *createVerificationRedisCommand) Execute(context context.Context, verification *entities.Verification) error {
	key := getKey(verification.AccountId)
	verificationsPtr, err := commands.GetValueIfExists[[]*entities.Verification](context, c.client, key)
	if err != nil {
		return err
	}

	var verifications []*entities.Verification
	if verificationsPtr == nil {
		verifications = make([]*entities.Verification, 0)
	} else {
		verifications = *verificationsPtr
	}

	verifications = append(verifications, verification)

	duration := verification.ExpirationTime.Sub(time.Now())
	err = commands.SetValue[[]*entities.Verification](context, c.client, key, verifications, duration)
	if err != nil {
		return err
	}

	return nil
}
