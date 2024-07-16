package usecases

import (
	"context"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	avatar2 "smartri_app/internal/entities/avatar_entities"
)

type updateAvatarUseCase struct {
	avatarRepo internal.IAvatarRepository
}

func NewInitOrUpdateAvatarUseCase(avatarRepo internal.IAvatarRepository) internal.IInitOrUpdateAvatarUseCase {
	return &updateAvatarUseCase{avatarRepo: avatarRepo}
}

func (uc *updateAvatarUseCase) InitOrUpdate(context context.Context, accountId int, request *requests.AvatarRequest) error {
	avatar, err := uc.avatarRepo.GetByAccountId(context, accountId)

	if err != nil {
		return err
	}

	hairColor := uc.getColor(request.HairColor)
	eyesColor := uc.getColor(request.EyesColor)
	skinColor := uc.getColor(request.SkinColor)

	newAvatar := avatar2.NewAvatar(accountId, request.HairId, hairColor, request.EyesId, eyesColor, request.ClothesId, request.ExpressionId, skinColor)

	if avatar == nil {
		err = uc.avatarRepo.Create(context, newAvatar)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = uc.avatarRepo.Update(context, accountId, newAvatar)
	return err
}

func (uc *updateAvatarUseCase) getColor(request *requests.ColorRequest) *avatar2.ColorRGBA {
	return avatar2.NewColorRGBA(request.R, request.G, request.B, request.A)
}
