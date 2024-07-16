package requests

import "time"

type AccountDevicesResponse struct {
	AccountId int               `json:"accountId"`
	Devices   []*DeviceResponse `json:"devices"`
}

type DeviceResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creationDate"`
}
