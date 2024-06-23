package entities

type User struct {
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	XP        int    `json:"XP"`
	AccountId int    `json:"account_id"`
}
