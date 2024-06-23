package requests

type UserAnswersRequest struct {
	Answers []Answer `json:"answers"`
}
type Answer struct {
	QuestionId int `json:"questionId"`
	AnswerId   int `json:"answerId"`
}
