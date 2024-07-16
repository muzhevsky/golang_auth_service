package usecases

import (
	"authorization/controllers/requests"
	"authorization/internal"
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
		deviceResponses[i] = requests.NewDeviceResponse(devices[i].Id, devices[i].Name, devices[i].SessionCreationTime)
	}

	return requests.NewAccountDevicesResponse(accountId, deviceResponses), nil
}
