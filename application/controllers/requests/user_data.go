package requests

import (
	"smartri_app/internal/entities/user_data"
)

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
	Skills    []*user_data.UserSkill `json:"skills"`
}

type AddSkillChangeRequest struct {
	Points  int `json:"points"`
	SkillId int `json:"skillId"`
}

type AvatarRequest struct {
	HairId       int          `json:"hairId"`
	HairColor    ColorRequest `json:"hairColor"`
	EyesId       int          `json:"eyesId"`
	EyesColor    ColorRequest `json:"eyesColor"`
	ClothesId    int          `json:"clothesId"`
	ExpressionId int          `json:"expressionId"`
	SkinColor    ColorRequest `json:"skinColor"`
}

type ColorRequest struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}
