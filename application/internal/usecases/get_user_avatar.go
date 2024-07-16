package usecases

import (
	"context"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	"smartri_app/internal/errs"
)

type getUserAvatarUseCase struct {
	repo internal.IAvatarRepository
}

func NewGetUserAvatarUseCase(repo internal.IAvatarRepository) internal.IGetUserAvatarUseCase {
	return &getUserAvatarUseCase{repo: repo}
}

func (uc *getUserAvatarUseCase) GetAvatar(context context.Context, accountId int) (*requests.AvatarRequest, error) {
	result, err := uc.repo.GetByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errs.UserDoesntHaveAnAvatarYetError
	}

	hairColorResponse := requests.NewColorRequest(result.HairColor.R, result.HairColor.G, result.HairColor.B, result.HairColor.A)
	eyesColorResponse := requests.NewColorRequest(result.EyesColor.R, result.EyesColor.G, result.EyesColor.B, result.EyesColor.A)
	skinColorResponse := requests.NewColorRequest(result.SkinColor.R, result.SkinColor.G, result.SkinColor.B, result.SkinColor.A)

	return requests.NewAvatarRequest(result.HairId, hairColorResponse, result.EyesId, eyesColorResponse, result.ClothesId, result.ExpressionId, skinColorResponse), nil
}
