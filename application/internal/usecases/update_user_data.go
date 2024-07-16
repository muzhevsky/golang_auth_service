package usecases

import (
	"context"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	"smartri_app/internal/entities/user_data"
)

type createOrUpdateUserDataUseCase struct {
	repo internal.IUserDataRepository
}

func NewAddOrUpdateUserDataUseCase(repo internal.IUserDataRepository) internal.IInitOrUpdateUserDataUseCase {
	return &createOrUpdateUserDataUseCase{repo: repo}
}

func (u *createOrUpdateUserDataUseCase) InitOrUpdate(context context.Context, request *requests.UserDataRequest, accountId int) (*requests.UserDataResponse, error) {
	existingUserData, err := u.repo.GetByAccountId(context, accountId)

	if err != nil {
		return nil, err
	}

	if existingUserData == nil {
		return u.create(context, request, accountId)
	}

	return u.update(context, request, existingUserData)

}

func (u *createOrUpdateUserDataUseCase) create(context context.Context, request *requests.UserDataRequest, accountId int) (*requests.UserDataResponse, error) {
	user, err := u.createUserData(request, accountId)
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(context, user)
	if err != nil {
		return nil, err
	}

	return requests.NewUserDataResponse(request.Nickname, request.Age, request.Gender, 0), nil
}

func (u *createOrUpdateUserDataUseCase) update(context context.Context, request *requests.UserDataRequest, user *user_data.UserData) (*requests.UserDataResponse, error) {
	user.Nickname = user_data.Nickname(request.Nickname)
	user.Age = user_data.Age(request.Age)
	user.Gender = user_data.Gender(request.Gender)

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	result, err := u.repo.Update(context, user)
	if err != nil {
		return nil, err
	}

	response := requests.NewUserDataResponse(request.Nickname, int(result.Age), string(result.Gender), int(result.XP))
	return response, nil
}

func (u *createOrUpdateUserDataUseCase) createUserData(data *requests.UserDataRequest, accountId int) (*user_data.UserData, error) {
	return user_data.NewUserData(
		user_data.Nickname(data.Nickname),
		user_data.Age(data.Age),
		user_data.Gender(data.Gender),
		user_data.XP(0),
		accountId)
}
