package repositories

import (
	"authorization/internal"
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type deviceRepo struct {
	insertDeviceCommand              datasources.IInsertDeviceCommand
	selectDevicesByAccountIdCommand  datasources.ISelectDevicesByAccountIdCommand
	selectDeviceByAccessTokenCommand datasources.ISelectDeviceByAccessTokenCommand
	selectDeviceByIdCommand          datasources.ISelectDeviceByIdCommand
	updateDeviceByIdCommand          datasources.IUpdateDeviceByAccessTokenCommand
	deleteDeviceByIdCommand          datasources.IDeleteDeviceByIdCommand
}

func NewDeviceRepo(
	insertDeviceCommand datasources.IInsertDeviceCommand,
	selectDevicesByAccountIdCommand datasources.ISelectDevicesByAccountIdCommand,
	selectDeviceByIdCommand datasources.ISelectDeviceByIdCommand,
	selectDeviceByAccessTokenCommand datasources.ISelectDeviceByAccessTokenCommand,
	updateDeviceByIdCommand datasources.IUpdateDeviceByAccessTokenCommand,
	deleteDeviceByIdCommand datasources.IDeleteDeviceByIdCommand) internal.IDeviceRepository {
	return &deviceRepo{
		insertDeviceCommand:              insertDeviceCommand,
		selectDevicesByAccountIdCommand:  selectDevicesByAccountIdCommand,
		selectDeviceByIdCommand:          selectDeviceByIdCommand,
		selectDeviceByAccessTokenCommand: selectDeviceByAccessTokenCommand,
		updateDeviceByIdCommand:          updateDeviceByIdCommand,
		deleteDeviceByIdCommand:          deleteDeviceByIdCommand}
}

func (repo *deviceRepo) Create(context context.Context, device *session.Device) error {
	return repo.insertDeviceCommand.Execute(context, device)
}

func (repo *deviceRepo) SelectByAccountId(context context.Context, accountId int) ([]*session.Device, error) {
	return repo.selectDevicesByAccountIdCommand.Execute(context, accountId)
}

func (repo *deviceRepo) DeleteById(context context.Context, deviceId int) error {
	return repo.deleteDeviceByIdCommand.Execute(context, deviceId)
}

func (repo *deviceRepo) SelectByAccessToken(context context.Context, accessToken string) (*session.Device, error) {
	return repo.selectDeviceByAccessTokenCommand.Execute(context, accessToken)
}

func (repo *deviceRepo) SelectById(context context.Context, id int) (*session.Device, error) {
	return repo.selectDeviceByIdCommand.Execute(context, id)
}

func (repo *deviceRepo) UpdateByAccessToken(context context.Context, accessToken string, device *session.Device) error {
	device, err := repo.selectDeviceByAccessTokenCommand.Execute(context, accessToken)
	if err != nil {
		return err
	}

	cpy := session.Device{
		Id:                  device.Id,
		AccountId:           device.AccountId,
		Name:                device.Name,
		SessionAccessToken:  device.SessionAccessToken,
		SessionCreationTime: device.SessionCreationTime,
	}

	err = repo.updateDeviceByIdCommand.Execute(context, accessToken, &cpy)
	return err
}
