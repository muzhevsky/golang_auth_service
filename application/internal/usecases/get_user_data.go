package usecases

import (
	"context"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	"smartri_app/internal/errs"
)

type getUserDataUseCase struct {
	repo internal.IUserDataRepository
}

func NewGetUserDataUseCase(repo internal.IUserDataRepository) internal.IGetUserDataUseCase {
	return &getUserDataUseCase{repo: repo}
}

func (c *getUserDataUseCase) GetUserData(context context.Context, accountId int) (*requests.UserDataResponse, error) {
	userData, err := c.repo.GetByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, errs.UserDataNotFoundError
	}

	return requests.NewUserDataResponse(
		string(userData.Nickname),
		int(userData.Age),
		string(userData.Gender),
		int(userData.XP)), nil
}
