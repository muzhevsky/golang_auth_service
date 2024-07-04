package usecases

import (
	"context"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	"smartri_app/internal/entities/user_data"
)

type addOrUpdateUserDataUseCase struct {
	repo internal.IUserDataRepository
}

func NewAddOrUpdateUserDataUseCase(repo internal.IUserDataRepository) internal.IInitOrUpdateUserDataUseCase {
	return &addOrUpdateUserDataUseCase{repo: repo}
}

func (u *addOrUpdateUserDataUseCase) InitOrUpdate(context context.Context, data *requests.AddUserDataRequest, accountId int) (*requests.UserDataResponse, error) {
	user, err := u.repo.GetByAccountId(context, accountId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		err = u.repo.Create(context, &user_data.UserData{
			Age:       data.Age,
			Gender:    data.Gender,
			XP:        0,
			AccountId: accountId,
		})
		if err != nil {
			return nil, err
		}
		return &requests.UserDataResponse{
			Age:    data.Age,
			Gender: data.Gender,
			XP:     0,
		}, nil
	}

	user.Age = data.Age
	user.Gender = data.Gender

	result, err := u.repo.Update(context, user)
	if err != nil {
		return nil, err
	}

	response := &requests.UserDataResponse{
		Age:    result.Age,
		Gender: result.Gender,
		XP:     result.XP,
	}
	return response, nil
}
