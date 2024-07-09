package session

import "time"

type Device struct {
	Identity            string    `json:"identity"`
	Name                string    `json:"name"`
	SessionAccessToken  string    `json:"session_access_token"`
	SessionCreationTime time.Time `json:"session_creation_time"`
}
