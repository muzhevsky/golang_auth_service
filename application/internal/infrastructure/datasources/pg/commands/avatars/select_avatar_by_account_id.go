package avatars

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	avatar2 "smartri_app/internal/entities/avatar_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAvatarByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectAvatarByAccountIdPGCommand(client *postgres.Client) datasources.ISelectAvatarByAccountIdCommand {
	return &selectAvatarByAccountIdPGCommand{client: client}
}

func (c *selectAvatarByAccountIdPGCommand) Execute(context context.Context, accountId int) (*avatar2.Avatar, error) {
	sql, args, err := query_builders.NewSelectAvatarByAccountIdQuery(&c.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	row := c.client.Pool.QueryRow(context, sql, args...)
	result := avatar2.Avatar{AccountId: accountId}
	hairColorInt := int32(0)
	eyesColorInt := int32(0)
	skinColorInt := int32(0)
	err = row.Scan(&result.HairId, &hairColorInt, &result.EyesId, &eyesColorInt, &result.ClothesId, &result.ExpressionId, &skinColorInt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	result.HairColor = avatar2.NewColorRGBAFromInt32(hairColorInt)
	result.EyesColor = avatar2.NewColorRGBAFromInt32(eyesColorInt)
	result.SkinColor = avatar2.NewColorRGBAFromInt32(skinColorInt)

	return &result, nil
}
