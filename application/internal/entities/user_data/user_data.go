package user_data

import "smartri_app/internal/errs"

type UserData struct {
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	XP        int    `json:"XP"`
	AccountId int    `json:"account_id"`
}

func NewUserData(age int, gender string, XP int, accountId int) (*UserData, error) {
	if age < 0 || age > 100 {
		return nil, errs.SomeErrorToDo
	}

	if gender != "m" && gender != "f" {
		return nil, errs.SomeErrorToDo
	}

	if XP < 0 {
		return nil, errs.SomeErrorToDo
	}

	return &UserData{Age: age, Gender: gender, XP: XP, AccountId: accountId}, nil
}
