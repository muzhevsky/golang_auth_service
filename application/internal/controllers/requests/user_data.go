package requests

import "smartri_app/internal/entities"

type AddUserDataRequest struct {
	Age    int    `json:"age" example:"11"`
	Gender string `json:"gender" example:"m"`
}

type UserDataResponse struct {
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	XP     int    `json:"xp"`
}

type UserSkillResponse struct {
	AccountId int                    `json:"accountId"`
	Skills    []*entities.UserSkills `json:"skills"`
}

type AddSkillChangeRequest struct {
	Points  int `json:"points"`
	SkillId int `json:"skillId"`
}
