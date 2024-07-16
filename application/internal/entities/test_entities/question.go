package test_entities

type Question struct {
	Id      int       `json:"id"`
	Text    string    `json:"text"`
	Answers []*Answer `json:"answers,omitempty"`
}
