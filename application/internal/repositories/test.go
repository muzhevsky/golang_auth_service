package repositories

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
)

type testRepository struct {
	questionsDataSource    datasources.IQuestionDataSource
	answersDataSource      datasources.IAnswerDataSource
	answerValuesDataSource datasources.IAnswerValuesDataSource
	userAnswersDataSource  datasources.IUserTestResultsDataSource
}

func NewTestRepository(
	questionsDataSource datasources.IQuestionDataSource,
	answersDataSource datasources.IAnswerDataSource,
	answerValuesDataSource datasources.IAnswerValuesDataSource,
	userAnswersDataSource datasources.IUserTestResultsDataSource) *testRepository {
	return &testRepository{
		questionsDataSource:    questionsDataSource,
		answersDataSource:      answersDataSource,
		answerValuesDataSource: answerValuesDataSource,
		userAnswersDataSource:  userAnswersDataSource,
	}
}

func (repo *testRepository) GetAllQuestions(context context.Context) ([]*entities.Question, error) {
	return repo.questionsDataSource.SelectAll(context)
}

func (repo *testRepository) GetAnswersForQuestion(context context.Context, question *entities.Question) ([]*entities.Answer, error) {
	return repo.answersDataSource.SelectByQuestionId(context, question.Id)
}

func (repo *testRepository) GetAllQuestionsWithAnswers(context context.Context) ([]*entities.Question, error) {
	questions, err := repo.questionsDataSource.SelectAll(context)
	if err != nil {
		return nil, err
	}

	for _, question := range questions {
		answers, err := repo.answersDataSource.SelectByQuestionId(context, question.Id)
		if err != nil {
			return nil, err
		}

		question.Answers = answers
	}

	return questions, nil
}

func (repo *testRepository) AddTestAnswers(context context.Context, results *entities.UserTestAnswers) error {
	err := repo.userAnswersDataSource.Insert(context, results)
	return err
}
