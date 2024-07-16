package requests

import (
	"smartri_app/internal/entities/skills_entities"
)

type UserAnswersRequest struct {
	Answers []Answer `json:"answers"`
}
type Answer struct {
	QuestionId int `json:"questionId"`
	AnswerId   int `json:"answerId"`
}
type UserAnswersResponse struct {
	AccountId int                          `json:"accountId"`
	Skills    []*skills_entities.UserSkill `json:"skills"`
	TotalExp  int                          `json:"totalExp"`
}
