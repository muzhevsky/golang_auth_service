package repositories

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/entities/skills"
	"smartri_app/internal/entities/test"
	"smartri_app/internal/entities/user_data"
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

func (repo *testRepository) GetAllQuestions(context context.Context) ([]*test.Question, error) {
	return repo.selectAllQuestions.Execute(context)
}

func (repo *testRepository) GetAnswersForQuestion(context context.Context, question *test.Question) ([]*test.Answer, error) {
	return repo.selectAnswersByQuestionId.Execute(context, question.Id)
}

func (repo *testRepository) GetAllQuestionsWithAnswers(context context.Context) ([]*test.Question, error) {
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
	answers *test.UserTestAnswers,
	changes []*skills.SkillChange,
	userSkills *skills.UserSkills,
	data *user_data.UserData) error {
	err := repo.insertUserAnswers.Execute(context, answers, changes, userSkills, data)
	return err
}

func (repo *testRepository) GetAnswerWithValues(context context.Context, answerId int) (*test.Answer, error) {
	values, err := repo.selectAnswerWithAnswerValuesById.Execute(context, answerId)
	if err != nil {
		return nil, err
	}

	answer, err := repo.selectAnswerById.Execute(context, answerId)
	if err != nil {
		return nil, err
	}
	answer.Values = values

	return answer, nil
}
