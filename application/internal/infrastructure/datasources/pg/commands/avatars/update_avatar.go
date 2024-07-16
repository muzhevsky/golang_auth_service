package avatars

import (
	"context"
	avatarpkg "smartri_app/internal/entities/avatar_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type updateAvatarByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewUpdateAvatarByAccountIdPGCommand(client *postgres.Client) datasources.IUpdateAvatarCommand {
	return &updateAvatarByAccountIdPGCommand{client: client}
}

func (c *updateAvatarByAccountIdPGCommand) Execute(context context.Context, accountId int, avatar *avatarpkg.Avatar) (*avatarpkg.Avatar, error) {
	sql, args, err := query_builders.NewUpdateAvatarByAccountIdQuery(&c.client.Builder, accountId, avatar)
	if err != nil {
		return nil, err
	}

	_, err = c.client.Pool.Exec(context, sql, args...)
	if err != nil {
		return nil, err
	}

	result := &avatarpkg.Avatar{
		AccountId:    accountId,
		HairId:       avatar.HairId,
		HairColor:    avatar.HairColor,
		EyesId:       avatar.EyesId,
		EyesColor:    avatar.EyesColor,
		ClothesId:    avatar.ClothesId,
		ExpressionId: avatar.ExpressionId,
		SkinColor:    avatar.SkinColor,
	}

	return result, nil
}
