package usecases

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/controllers/requests"
	entities2 "smartri_app/internal/entities"
)

type addUserAnswers struct {
	repo internal.ITestRepository
}

func NewAddUserAnswers(repo internal.ITestRepository) *addUserAnswers {
	return &addUserAnswers{repo: repo}
}

func (a *addUserAnswers) Add(context context.Context, answers *requests.UserAnswersRequest, accountId int) error {
	entityAnswers := entities2.UserTestAnswers{accountId, make([]entities2.UserTestAnswer, 0)}

	for i := range answers.Answers {
		entityAnswers.Answers = append(entityAnswers.Answers, entities2.UserTestAnswer{
			QuestionId: answers.Answers[i].QuestionId,
			AnswerId:   answers.Answers[i].AnswerId,
		})
	}

	err := a.repo.AddTestAnswers(context, &entityAnswers)
	return err
}
