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
	devices, err := uc.deviceRepository.SelectDevicesByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	deviceResponses := make([]*requests.DeviceResponse, len(devices))
	for i := 0; i < len(devices); i++ {
		deviceResponses[i] = &requests.DeviceResponse{
			Id:           devices[i].Id,
			Name:         devices[i].Name,
			CreationDate: devices[i].SessionCreationTime,
		}
	}

	return &requests.AccountDevicesResponse{
		AccountId: accountId,
		Devices:   deviceResponses,
	}, nil
}
