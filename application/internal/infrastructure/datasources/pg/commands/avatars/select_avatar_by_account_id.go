package avatars

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/user_data/avatar"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAvatarByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectAvatarByAccountIdPGCommand(client *postgres.Client) *selectAvatarByAccountIdPGCommand {
	return &selectAvatarByAccountIdPGCommand{client: client}
}

func (c *selectAvatarByAccountIdPGCommand) Execute(context context.Context, accountId int) (*avatar.Avatar, error) {
	sql, args, err := query_builders.NewSelectAvatarByAccountIdQuery(&c.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	row := c.client.Pool.QueryRow(context, sql, args...)
	result := avatar.Avatar{AccountId: accountId}
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

	result.HairColor = avatar.NewColorRGBAFromInt32(hairColorInt)
	result.EyesColor = avatar.NewColorRGBAFromInt32(eyesColorInt)
	result.SkinColor = avatar.NewColorRGBAFromInt32(skinColorInt)

	return &result, nil
}
