package usecases

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/entities"
)

type addUserData struct {
	repo internal.IUserDataRepository
}

func NewAddUserData(repo internal.IUserDataRepository) *addUserData {
	return &addUserData{repo: repo}
}

func (a *addUserData) Add(context context.Context, details *requests.UserDataRequest, accountId int) (*requests.UserDataResponse, error) {
	user := &entities.User{
		Age:       details.Age,
		Gender:    details.Gender,
		XP:        0,
		AccountId: accountId,
	}

	err := a.repo.AddOrUpdate(context, user)
	if err != nil {
		return nil, err
	}

	user, err = a.repo.GetByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	result := &requests.UserDataResponse{
		Age:    user.Age,
		Gender: user.Gender,
		XP:     user.XP,
	}
	return result, nil
}
