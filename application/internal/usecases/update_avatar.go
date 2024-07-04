package usecases

import (
	"context"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	avatarpkg "smartri_app/internal/entities/user_data/avatar"
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

	hairColor := uc.getHairColor(request)
	eyeColor := uc.getEyeColor(request)
	skinColor := uc.getSkinColor(request)

	newAvatar := &avatarpkg.Avatar{
		AccountId:    accountId,
		HairId:       request.HairId,
		HairColor:    hairColor,
		EyesId:       request.EyesId,
		EyesColor:    eyeColor,
		ClothesId:    request.ClothesId,
		ExpressionId: request.ExpressionId,
		SkinColor:    skinColor,
	}

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

func (uc *updateAvatarUseCase) getHairColor(request *requests.AvatarRequest) *avatarpkg.ColorRGBA {
	return avatarpkg.NewColorRGBA(
		request.HairColor.R,
		request.HairColor.G,
		request.HairColor.B,
		request.HairColor.A)
}

func (uc *updateAvatarUseCase) getEyeColor(request *requests.AvatarRequest) *avatarpkg.ColorRGBA {
	return avatarpkg.NewColorRGBA(
		request.EyesColor.R,
		request.EyesColor.G,
		request.EyesColor.B,
		request.EyesColor.A)
}

func (uc *updateAvatarUseCase) getSkinColor(request *requests.AvatarRequest) *avatarpkg.ColorRGBA {
	return avatarpkg.NewColorRGBA(
		request.SkinColor.R,
		request.SkinColor.G,
		request.SkinColor.B,
		request.SkinColor.A)
}
