package repositories

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
)

type userRepository struct {
	selectUserDataByAccountIdCommand        datasources.ISelectUserDataByAccountIdCommand
	updateUserDataCommand                   datasources.IUpdateUserDataByAccountIdCommand
	insertUserDataCommand                   datasources.IInsertUserDataCommand
	checkIfUserHasAnswersByAccountIdCommand datasources.ICheckIfUserHasAnswersByAccountIdCommand
	applySkillChangesByAccountIdCommand     datasources.IApplySkillChangesByAccountIdCommand
}

func NewUserRepository(
	selectUserDataByAccountIdCommand datasources.ISelectUserDataByAccountIdCommand,
	updateUserDataCommand datasources.IUpdateUserDataByAccountIdCommand,
	insertUserDataCommand datasources.IInsertUserDataCommand,
	checkIfUserHasAnswersCommand datasources.ICheckIfUserHasAnswersByAccountIdCommand,
	applySkillChangesByAccountIdCommand datasources.IApplySkillChangesByAccountIdCommand) internal.IUserDataRepository {
	return &userRepository{
		selectUserDataByAccountIdCommand:        selectUserDataByAccountIdCommand,
		updateUserDataCommand:                   updateUserDataCommand,
		insertUserDataCommand:                   insertUserDataCommand,
		checkIfUserHasAnswersByAccountIdCommand: checkIfUserHasAnswersCommand,
		applySkillChangesByAccountIdCommand:     applySkillChangesByAccountIdCommand}
}

func (u *userRepository) GetUserDataByAccountId(context context.Context, accountId int) (*entities.UserData, error) {
	return u.selectUserDataByAccountIdCommand.Execute(context, accountId)
}

func (u *userRepository) UpdateUserData(context context.Context, details *entities.UserData) (*entities.UserData, error) {
	updated, err := u.updateUserDataCommand.Execute(context, details)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (u *userRepository) AddUserData(context context.Context, details *entities.UserData) error {
	err := u.insertUserDataCommand.Execute(context, details)
	return err
}

func (u *userRepository) CheckUserHasAnswers(context context.Context, accountId int) (bool, error) {
	return u.checkIfUserHasAnswersByAccountIdCommand.Execute(context, accountId)
}

func (u *userRepository) ApplySkillChangesByAccountId(context context.Context, userSkills *entities.UserSkills, userData *entities.UserData, change *entities.SkillChange) error {
	return u.applySkillChangesByAccountIdCommand.Execute(context, userSkills, userData, change)
}
