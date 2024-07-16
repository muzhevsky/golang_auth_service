package user_data

type UserData struct {
	AccountId int      `json:"account_id"`
	Nickname  Nickname `json:"nickname"`
	Age       Age      `json:"age"`
	Gender    Gender   `json:"gender"`
	XP        XP       `json:"XP"`
}

func NewUserData(nickname Nickname, age Age, gender Gender, XP XP, accountId int) (*UserData, error) {
	data := &UserData{Nickname: nickname, Age: age, Gender: gender, XP: XP, AccountId: accountId}
	err := data.Validate()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ud *UserData) Validate() error {
	err := ud.Nickname.Validate()
	if err != nil {
		return err
	}
	err = ud.Age.Validate()
	if err != nil {
		return err
	}
	err = ud.Gender.Validate()
	if err != nil {
		return err
	}
	err = ud.XP.Validate()

	return err
}
