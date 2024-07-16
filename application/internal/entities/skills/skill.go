package skills

import "time"

type Skill struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type SkillNormalization struct {
	SkillId int `json:"skillId"`
	Min     int `json:"min"`
	Max     int `json:"max"`
}

type SkillChange struct {
	Id        int       `json:"id"`
	AccountId int       `json:"accountId"`
	SkillId   int       `json:"skillId"`
	ActionId  int       `json:"actionId"`
	Date      time.Time `json:"date"`
	Points    int       `json:"points"`
}
