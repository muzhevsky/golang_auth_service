package entities

type UserData struct {
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	XP        int    `json:"XP"`
	AccountId int    `json:"account_id"`
}

type UserSkills struct {
	AccountId int `json:"account_id,omitempty"`
	SkillId   int `json:"skill_id"`
	Xp        int `json:"points"`
}
