package internal

import (
	"context"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/entities"
)

type (
	ITestRepository interface {
		GetAllQuestions(context context.Context) ([]*entities.Question, error)
		GetAnswersForQuestion(context context.Context, question *entities.Question) ([]*entities.Answer, error)
		GetAllQuestionsWithAnswers(context context.Context) ([]*entities.Question, error)
		AddTestAnswers(context context.Context, results *entities.UserTestAnswers) error
	}

	IUserDataRepository interface {
		GetByAccountId(context context.Context, accountId int) (*entities.User, error)
		AddOrUpdate(context context.Context, details *entities.User) error
	}

	IAddUserDataUseCase interface {
		Add(context context.Context, details *requests.UserDataRequest, accountId int) (*requests.UserDataResponse, error)
	}

	IAddUserTestAnswersUseCase interface {
		Add(context context.Context, answers *requests.UserAnswersRequest, accountId int) error
	}
)
