package repositories

import (
	"context"
	"fmt"
	"smartri_app/internal"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/errs"
	"smartri_app/internal/infrastructure/datasources"
)

type testRepository struct {
	selectAllQuestions               datasources.ISelectAllQuestionsCommand
	selectAnswerById                 datasources.ISelectAnswerByIdCommand
	selectAnswersByQuestionId        datasources.ISelectAnswersByQuestionIdCommand
	selectAnswerWithAnswerValuesById datasources.ISelectAnswerValuesByAnswerIdCommand
	insertUserAnswers                datasources.IInsertUserTestResultsCommand
}

func NewTestRepository(
	questionsDataSource datasources.ISelectAllQuestionsCommand,
	answersDataSource datasources.ISelectAnswerByIdCommand,
	selectAnswersByQuestionId datasources.ISelectAnswersByQuestionIdCommand,
	answerValuesDataSource datasources.ISelectAnswerValuesByAnswerIdCommand,
	userAnswersDataSource datasources.IInsertUserTestResultsCommand) internal.ITestRepository {
	return &testRepository{
		selectAllQuestions:               questionsDataSource,
		selectAnswerById:                 answersDataSource,
		selectAnswersByQuestionId:        selectAnswersByQuestionId,
		selectAnswerWithAnswerValuesById: answerValuesDataSource,
		insertUserAnswers:                userAnswersDataSource,
	}
}

func (repo *testRepository) GetAllQuestions(context context.Context) ([]*test_entities.Question, error) {
	return repo.selectAllQuestions.Execute(context)
}

func (repo *testRepository) GetAnswersForQuestion(context context.Context, question *test_entities.Question) ([]*test_entities.Answer, error) {
	return repo.selectAnswersByQuestionId.Execute(context, question.Id)
}

func (repo *testRepository) GetAllQuestionsWithAnswers(context context.Context) ([]*test_entities.Question, error) {
	questions, err := repo.selectAllQuestions.Execute(context)
	if err != nil {
		return nil, err
	}

	for _, question := range questions {
		answers, err := repo.selectAnswersByQuestionId.Execute(context, question.Id)
		if err != nil {
			return nil, err
		}

		question.Answers = answers
	}

	return questions, nil
}

func (repo *testRepository) AddUserAnswersWithSkillChanges(
	context context.Context,
	answers *test_entities.UserTestAnswers,
	changes []*skills_entities.SkillChange,
	userSkills *skills_entities.UserSkills,
	data *user_data_entities.UserData) error {
	err := repo.insertUserAnswers.Execute(context, answers, changes, userSkills, data)
	return err
}

func (repo *testRepository) GetAnswerWithValues(context context.Context, answerId int) (*test_entities.Answer, error) {
	values, err := repo.selectAnswerWithAnswerValuesById.Execute(context, answerId)
	if err != nil {
		return nil, err
	}

	answer, err := repo.selectAnswerById.Execute(context, answerId)
	if err != nil {
		return nil, err
	}
	if answer == nil {
		return nil, fmt.Errorf("%w. answer not found", errs.EntityNotFoundError)
	}
	answer.Values = values

	return answer, nil
}
