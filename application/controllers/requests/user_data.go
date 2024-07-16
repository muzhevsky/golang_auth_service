package requests

import (
	"smartri_app/internal/entities/skills_entities"
)

type UserDataRequest struct {
	Nickname string `json:"nickname" example:"SlimShady123"`
	Age      int    `json:"age" example:"11"`
	Gender   string `json:"gender" example:"m"`
}

type UserDataResponse struct {
	Nickname string `json:"nickname" example:"SlimShady123"`
	Age      int    `json:"age" example:"11"`
	Gender   string `json:"gender" example:"m"`
	XP       int    `json:"xp" example:"862"`
}

func NewUserDataResponse(nickname string, age int, gender string, xp int) *UserDataResponse {
	return &UserDataResponse{Nickname: nickname, Age: age, Gender: gender, XP: xp}
}

type UserSkillResponse struct {
	AccountId int                          `json:"accountId"`
	Skills    []*skills_entities.UserSkill `json:"skills"`
}

type AddSkillChangeRequest struct {
	Points  int `json:"points" example:"10"`
	SkillId int `json:"skillId" example:"5"`
}

type AvatarRequest struct {
	HairId       int           `json:"hairId"`
	HairColor    *ColorRequest `json:"hairColor"`
	EyesId       int           `json:"eyesId"`
	EyesColor    *ColorRequest `json:"eyesColor"`
	ClothesId    int           `json:"clothesId"`
	ExpressionId int           `json:"expressionId"`
	SkinColor    *ColorRequest `json:"skinColor"`
}

func NewAvatarRequest(hairId int, hairColor *ColorRequest, eyesId int, eyesColor *ColorRequest, clothesId int, expressionId int, skinColor *ColorRequest) *AvatarRequest {
	return &AvatarRequest{HairId: hairId, HairColor: hairColor, EyesId: eyesId, EyesColor: eyesColor, ClothesId: clothesId, ExpressionId: expressionId, SkinColor: skinColor}
}

type ColorRequest struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

func NewColorRequest(r uint8, g uint8, b uint8, a uint8) *ColorRequest {
	return &ColorRequest{R: r, G: g, B: b, A: a}
}
