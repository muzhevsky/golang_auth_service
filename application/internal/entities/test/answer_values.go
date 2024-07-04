package test

type AnswerValue struct {
	Id       int `json:"id"`
	AnswerId int `json:"answer_id"`
	SkillId  int `json:"skill_id"`
	Points   int `json:"points"`
}
