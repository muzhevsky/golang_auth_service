package verification

import (
	verificationpkg "authorization/internal/entities/verification_entities"
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

func (c *createVerificationRedisCommand) Execute(context context.Context, verification *verificationpkg.Verification) error {
	key := getKey(verification.AccountId)
	verificationsPtr, err := commands.GetValueOrNil[[]*verificationpkg.Verification](context, c.client, key)
	if err != nil {
		return err
	}

	var verifications []*verificationpkg.Verification
	if verificationsPtr == nil {
		verifications = make([]*verificationpkg.Verification, 0)
	} else {
		verifications = *verificationsPtr
	}

	verifications = append(verifications, verification)

	duration := verification.ExpirationTime.Sub(time.Now())
	err = commands.SetValue(context, c.client, key, verifications, duration)
	if err != nil {
		return err
	}

	return nil
}
