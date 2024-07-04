package query_builders

import (
	"github.com/Masterminds/squirrel"
	"smartri_app/internal/entities/user_data/avatar"
)

const avatarsTableName = "avatars"
const avatarsAccountIdFieldName = "account_id"
const avatarsHairIdFieldName = "hair_id"
const avatarsHairColorFieldName = "hair_color"
const avatarsEyesIdFieldName = "eyes_id"
const avatarsEyesColorFieldName = "eyes_color"
const avatarsClothesIdFieldName = "clothes_id"
const avatarsExpressionIdFieldName = "expression_id"
const avatarsSkinColorFieldName = "skin_color"

func NewSelectAvatarByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int) (string, []any, error) {
	return builder.
		Select(
			avatarsHairIdFieldName,
			avatarsHairColorFieldName,
			avatarsEyesIdFieldName,
			avatarsEyesColorFieldName,
			avatarsClothesIdFieldName,
			avatarsExpressionIdFieldName,
			avatarsSkinColorFieldName).
		From(avatarsTableName).
		Where(squirrel.Eq{avatarsAccountIdFieldName: accountId}).
		ToSql()
}

func NewInsertAvatarQuery(builder *squirrel.StatementBuilderType, avatar *avatar.Avatar) (string, []any, error) {
	return builder.
		Insert(avatarsTableName).
		Columns(
			avatarsAccountIdFieldName,
			avatarsHairIdFieldName,
			avatarsHairColorFieldName,
			avatarsEyesIdFieldName,
			avatarsEyesColorFieldName,
			avatarsClothesIdFieldName,
			avatarsExpressionIdFieldName,
			avatarsSkinColorFieldName).
		Values(
			avatar.AccountId,
			avatar.HairId,
			avatar.HairColor.ToInt32(),
			avatar.EyesId,
			avatar.EyesColor.ToInt32(),
			avatar.ClothesId,
			avatar.ExpressionId,
			avatar.SkinColor.ToInt32()).
		ToSql()
}

func NewUpdateAvatarByAccountIdQuery(builder *squirrel.StatementBuilderType, accountId int, avatar *avatar.Avatar) (string, []any, error) {
	return builder.
		Update(avatarsTableName).
		Set(avatarsHairIdFieldName, avatar.HairId).
		Set(avatarsHairColorFieldName, avatar.HairColor.ToInt32()).
		Set(avatarsEyesIdFieldName, avatar.EyesId).
		Set(avatarsEyesColorFieldName, avatar.EyesColor.ToInt32()).
		Set(avatarsClothesIdFieldName, avatar.ClothesId).
		Set(avatarsExpressionIdFieldName, avatar.ExpressionId).
		Set(avatarsSkinColorFieldName, avatar.SkinColor.ToInt32()).
		Where(squirrel.Eq{avatarsAccountIdFieldName: accountId}).
		ToSql()
}
