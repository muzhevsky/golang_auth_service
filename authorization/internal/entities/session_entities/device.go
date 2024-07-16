package session_entities

import "time"

type Device struct {
	Id                  int       `json:"id"`
	AccountId           int       `json:"accountId"`
	Name                string    `json:"name"`
	SessionAccessToken  string    `json:"sessionAccessToken"`
	SessionCreationTime time.Time `json:"sessionCreationTime"`
}

func NewDevice(accountId int, name string, sessionAccessToken string) *Device {
	return &Device{AccountId: accountId, Name: name, SessionAccessToken: sessionAccessToken, SessionCreationTime: time.Now()}
}
