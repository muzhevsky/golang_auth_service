package user_data

import (
	"context"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type updateUserDataByAccountIdPGCommand struct {
	client        *postgres.Client
	selectCommand datasources.ISelectUserDataByAccountIdCommand
}

func NewUpdateUserDataByAccountIdPGCommand(client *postgres.Client, selectCommand datasources.ISelectUserDataByAccountIdCommand) datasources.IUpdateUserDataByAccountIdCommand {
	return &updateUserDataByAccountIdPGCommand{
		client:        client,
		selectCommand: selectCommand,
	}
}

func (u *updateUserDataByAccountIdPGCommand) Execute(context context.Context, data *user_data_entities.UserData) (*user_data_entities.UserData, error) {
	user, err := u.selectCommand.Execute(context, data.AccountId)
	if err != nil || user == nil {
		return nil, err
	}

	sql, args, err := query_builders.NewUpdateUserDataByAccountIdQuery(&u.client.Builder, data)

	_, err = u.client.Pool.Exec(context, sql, args...)

	return data, err
}
