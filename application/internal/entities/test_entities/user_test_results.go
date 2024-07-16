package test_entities

type UserTestAnswers struct {
	AccountId int
	Answers   []UserTestAnswer
}

type UserTestAnswer struct {
	Id         int
	AccountId  int
	QuestionId int
	AnswerId   int
}
