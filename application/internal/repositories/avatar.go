package repositories

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/entities/user_data/avatar"
	"smartri_app/internal/infrastructure/datasources"
)

type avatarRepository struct {
	selectAvatarByAccountIdCommand datasources.ISelectAvatarByAccountIdCommand
	createAvatarCommand            datasources.IInsertAvatarCommand
	updateAvatarCommand            datasources.IUpdateAvatarCommand
}

func NewAvatarRepository(
	selectAvatarByAccountIdCommand datasources.ISelectAvatarByAccountIdCommand,
	createAvatarCommand datasources.IInsertAvatarCommand,
	updateAvatarCommand datasources.IUpdateAvatarCommand) internal.IAvatarRepository {
	return &avatarRepository{selectAvatarByAccountIdCommand: selectAvatarByAccountIdCommand, createAvatarCommand: createAvatarCommand, updateAvatarCommand: updateAvatarCommand}
}

func (a *avatarRepository) GetByAccountId(context context.Context, accountId int) (*avatar.Avatar, error) {
	return a.selectAvatarByAccountIdCommand.Execute(context, accountId)
}

func (a *avatarRepository) Create(context context.Context, avatar *avatar.Avatar) error {
	return a.createAvatarCommand.Execute(context, avatar)
}

func (a *avatarRepository) Update(context context.Context, accountId int, avatar *avatar.Avatar) (*avatar.Avatar, error) {
	return a.updateAvatarCommand.Execute(context, accountId, avatar)
}
