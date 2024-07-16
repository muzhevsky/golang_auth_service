package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	"authorization/internal/entities/session_entities"
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
	devicesToRemove := make([]*session_entities.Device, 0)
	ids := request.Ids
	userDevices, err := c.deviceRepo.SelectByAccountId(context, accountId)
	if err != nil {
		return err
	}

	for i := 0; i < len(ids); i++ { // todo maybe it worth it to move this loop to some entity
		index := slices.IndexFunc[[]*session_entities.Device](userDevices, func(d *session_entities.Device) bool {
			return d.Id == ids[i]
		})

		if index == -1 {
			continue
		}
		dev := userDevices[index]
		devicesToRemove = append(devicesToRemove, dev)
	}

	return c.removeDevicesAndSessions(context, devicesToRemove)
}

func (c *closeSessionsUseCase) removeDevicesAndSessions(context context.Context, devicesToRemove []*session_entities.Device) error {
	for i := 0; i < len(devicesToRemove); i++ {
		err := c.sessionRepo.DeleteByAccessToken(context, devicesToRemove[i].SessionAccessToken)
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
