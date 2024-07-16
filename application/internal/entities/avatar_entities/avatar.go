package avatar_entities

type Avatar struct {
	AccountId    int        `json:"account_id"`
	HairId       int        `json:"hairId"`
	HairColor    *ColorRGBA `json:"hairColor"`
	EyesId       int        `json:"eyesId"`
	EyesColor    *ColorRGBA `json:"eyesColor"`
	ClothesId    int        `json:"clothesId"`
	ExpressionId int        `json:"expressionId"`
	SkinColor    *ColorRGBA `json:"skinColor"`
}

func NewAvatar(
	accountId int,
	hairId int,
	hairColor *ColorRGBA,
	eyesId int,
	eyesColor *ColorRGBA,
	clothesId int,
	expressionId int,
	skinColor *ColorRGBA) *Avatar {
	return &Avatar{
		AccountId:    accountId,
		HairId:       hairId,
		HairColor:    hairColor,
		EyesId:       eyesId,
		EyesColor:    eyesColor,
		ClothesId:    clothesId,
		ExpressionId: expressionId,
		SkinColor:    skinColor}
}
