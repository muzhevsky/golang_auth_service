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

	return &requests.AvatarRequest{
		HairId: result.HairId,
		HairColor: requests.ColorRequest{
			R: result.HairColor.R,
			G: result.HairColor.G,
			B: result.HairColor.B,
			A: result.HairColor.A,
		},
		EyesId: result.EyesId,
		EyesColor: requests.ColorRequest{
			R: result.EyesColor.R,
			G: result.EyesColor.G,
			B: result.EyesColor.B,
			A: result.EyesColor.A,
		},
		ClothesId:    result.ClothesId,
		ExpressionId: result.ExpressionId,
		SkinColor: requests.ColorRequest{
			R: result.SkinColor.R,
			G: result.SkinColor.G,
			B: result.SkinColor.B,
			A: result.SkinColor.A,
		}}, nil
}
