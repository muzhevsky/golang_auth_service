package test

type Question struct {
	Id      int       `json:"id"`
	Text    string    `json:"text"`
	Answers []*Answer `json:"answers,omitempty"`
}
