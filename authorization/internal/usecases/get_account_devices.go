package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
	"authorization/internal/entities/session"
	"context"
)

type getAccountDevicesUseCase struct {
	deviceRepository internal.IDeviceRepository
}

func NewGetAccountDevicesUseCase(deviceRepository internal.IDeviceRepository) internal.IGetAccountDevicesUseCase {
	return &getAccountDevicesUseCase{deviceRepository: deviceRepository}
}

func (uc *getAccountDevicesUseCase) GetAccountDevices(context context.Context, accountId int) (*requests.AccountDevicesResponse, error) {
	devices, err := uc.deviceRepository.SelectByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	deviceResponses := make([]*requests.DeviceResponse, len(devices))
	for i := 0; i < len(devices); i++ {
		deviceResponses[i] = uc.createDeviceResponse(devices[i])
	}

	return uc.createResponse(accountId, deviceResponses), nil
}

func (uc *getAccountDevicesUseCase) createDeviceResponse(device *session.Device) *requests.DeviceResponse {
	return &requests.DeviceResponse{
		Id:           device.Id,
		Name:         device.Name,
		CreationDate: device.SessionCreationTime,
	}
}

func (uc *getAccountDevicesUseCase) createResponse(accountId int, deviceResponses []*requests.DeviceResponse) *requests.AccountDevicesResponse {
	return &requests.AccountDevicesResponse{
		AccountId: accountId,
		Devices:   deviceResponses,
	}
}
