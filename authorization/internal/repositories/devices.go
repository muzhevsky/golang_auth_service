package repositories

import (
	"authorization/internal"
	"authorization/internal/entities/session"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type deviceRepo struct {
	selectDevicesByAccountIdCommand datasources.ISelectDevicesByAccountIdCommand
	selectDeviceByIdCommand         datasources.ISelectDeviceByIdCommand
	deleteDeviceByIdCommand         datasources.IDeleteDeviceByIdCommand
	deleteSessionByIdCommand        datasources.IDeleteSessionByAccessTokenCommand
}

func NewDeviceRepo(
	selectDevicesByAccountIdCommand datasources.ISelectDevicesByAccountIdCommand,
	selectDeviceByIdCommand datasources.ISelectDeviceByIdCommand,
	deleteDeviceByIdCommand datasources.IDeleteDeviceByIdCommand,
	deleteSessionByIdCommand datasources.IDeleteSessionByAccessTokenCommand) internal.IDeviceRepository {
	return &deviceRepo{selectDevicesByAccountIdCommand: selectDevicesByAccountIdCommand, selectDeviceByIdCommand: selectDeviceByIdCommand, deleteDeviceByIdCommand: deleteDeviceByIdCommand, deleteSessionByIdCommand: deleteSessionByIdCommand}
}

func (repo *deviceRepo) SelectDevicesByAccountId(context context.Context, accountId int) ([]*session.Device, error) {
	return repo.selectDevicesByAccountIdCommand.Execute(context, accountId)
}

func (repo *deviceRepo) DeleteDeviceById(context context.Context, deviceId int) error {
	device, err := repo.selectDeviceByIdCommand.Execute(context, deviceId)
	if err != nil {
		return err
	}

	repo.deleteSessionByIdCommand.Execute(context, device.SessionAccessToken)

	return repo.DeleteDeviceById(context, deviceId)
}
