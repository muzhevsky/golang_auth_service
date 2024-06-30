package requests

import "smartri_app/internal/entities"

type UserAnswersRequest struct {
	Answers []Answer `json:"answers"`
}
type Answer struct {
	QuestionId int `json:"questionId"`
	AnswerId   int `json:"answerId"`
}
type UserAnswersResponse struct {
	AccountId int                    `json:"accountId"`
	Skills    []*entities.UserSkills `json:"skills"`
	TotalExp  int                    `json:"totalExp"`
}
