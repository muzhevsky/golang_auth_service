package requests

import "time"

type AccountDevicesResponse struct {
	AccountId int               `json:"accountId"`
	Devices   []*DeviceResponse `json:"devices"`
}

func NewAccountDevicesResponse(accountId int, devices []*DeviceResponse) *AccountDevicesResponse {
	return &AccountDevicesResponse{AccountId: accountId, Devices: devices}
}

type DeviceResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creationDate"`
}

func NewDeviceResponse(id int, name string, creationDate time.Time) *DeviceResponse {
	return &DeviceResponse{Id: id, Name: name, CreationDate: creationDate}
}
