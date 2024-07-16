package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	"authorization/internal/entities/session"
	"context"
	"slices"
)

type closeSessionsUseCase struct {
	deviceRepo  internal.IDeviceRepository
	sessionRepo internal.ISessionRepository
}

func NewCloseSessionsUseCase(deviceRepo internal.IDeviceRepository, sessionRepo internal.ISessionRepository) internal.ICloseSessionsByIdsUseCase {
	return &closeSessionsUseCase{deviceRepo: deviceRepo, sessionRepo: sessionRepo}
}

func (c *closeSessionsUseCase) CloseSessionsByIds(context context.Context, accountId int, request *requests.CloseSessionsRequest) error {
	devicesToRemove := make([]*session.Device, 0)
	ids := request.Ids
	userDevices, err := c.deviceRepo.SelectByAccountId(context, accountId)
	if err != nil {
		return err
	}

	for i := 0; i < len(ids); i++ {
		index := slices.IndexFunc[[]*session.Device](userDevices, func(d *session.Device) bool {
			return d.Id == ids[i]
		})

		if index == -1 {
			continue
		}
		dev := userDevices[index]
		devicesToRemove = append(devicesToRemove, dev)
	}

	for i := 0; i < len(devicesToRemove); i++ {
		err = c.sessionRepo.DeleteByAccessToken(context, devicesToRemove[i].SessionAccessToken)
		if err != nil {
			return err
		}
		err = c.deviceRepo.DeleteById(context, devicesToRemove[i].Id)
		if err != nil {
			return err
		}
	}

	return nil
}
