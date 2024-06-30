package usecases

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/entities"
)

type addOrUpdateUserDataUseCase struct {
	repo internal.IUserDataRepository
}

func NewAddOrUpdateUserDataUseCase(repo internal.IUserDataRepository) internal.IAddOrUpdateUserDataUseCase {
	return &addOrUpdateUserDataUseCase{repo: repo}
}

func (u *addOrUpdateUserDataUseCase) AddOrUpdate(context context.Context, data *requests.AddUserDataRequest, accountId int) (*requests.UserDataResponse, error) {
	user, err := u.repo.GetByAccountId(context, accountId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		err = u.repo.Add(context, &entities.UserData{
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
