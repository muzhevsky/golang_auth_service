package test

type UserTestAnswers struct {
	AccountId int
	Answers   []UserTestAnswer
}

type UserTestAnswer struct {
	QuestionId int
	AnswerId   int
}
