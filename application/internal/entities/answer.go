package entities

type Answer struct {
	Id         int            `json:"id"`
	QuestionId int            `json:"question_id"`
	Text       string         `json:"text"`
	Values     []*AnswerValue `json:"values,omitempty"`
}
