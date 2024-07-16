package session

import "time"

type Device struct {
	Id                  int       `json:"id"`
	AccountId           int       `json:"accountId"`
	Name                string    `json:"name"`
	SessionAccessToken  string    `json:"sessionAccessToken"`
	SessionCreationTime time.Time `json:"sessionCreationTime"`
}
