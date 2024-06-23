package datasources

import (
	"context"
	"smartri_app/internal/entities"
)

type (
	IQuestionDataSource interface {
		SelectAll(context context.Context) ([]*entities.Question, error)
		SelectById(context context.Context, id int) (*entities.Question, error)
	}

	IAnswerDataSource interface {
		SelectById(context context.Context, id int) (*entities.Answer, error)
		SelectByQuestionId(context context.Context, questionId int) ([]*entities.Answer, error)
	}

	IAnswerValuesDataSource interface {
		SelectById(context context.Context, id int) (*entities.AnswerValue, error)
		SelectByAnswerId(context context.Context, answerId int) ([]*entities.AnswerValue, error)
	}

	IUserDataSource interface {
		SelectByAccountId(context context.Context, id int) (*entities.User, error)
		Insert(context context.Context, user *entities.User) error
		UpdateByAccountId(context context.Context, accountId int, newUser *entities.User) error
	}

	IUserTestResultsDataSource interface {
		Insert(context context.Context, results *entities.UserTestAnswers) error
		//SelectByAccountId(context context.Context, id int) (*entities, error) // TODO
	}
)
