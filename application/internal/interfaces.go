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
		GetAnswerWithValues(context context.Context, answerId int) (*entities.Answer, error)
		GetAllQuestionsWithAnswers(context context.Context) ([]*entities.Question, error)
		AddUserTestResults(context context.Context, results *entities.UserTestAnswers, skills []*entities.SkillChange, userSkills []*entities.UserSkills, data *entities.UserData) error
	}

	ISkillRepository interface {
		GetAllSkills(context context.Context) ([]*entities.Skill, error)
		GetSkillNormalizationBySkillId(context context.Context, skillId int) (*entities.SkillNormalization, error)
		GetSkillsByAccountId(context context.Context, accountId int) ([]*entities.UserSkills, error)
	}

	IUserDataRepository interface {
		GetDataByAccountId(context context.Context, accountId int) (*entities.UserData, error)
		ApplySkillChangesByAccountId(context context.Context, userSkill *entities.UserSkills, userData *entities.UserData, change *entities.SkillChange) error
		CheckUserHasAnswers(context context.Context, accountId int) (bool, error)
		AddUserData(context context.Context, userData *entities.UserData) error
		UpdateUserData(context context.Context, details *entities.UserData) (*entities.UserData, error)
	}

	IAddOrUpdateUserDataUseCase interface {
		AddOrUpdate(context context.Context, data *requests.AddUserDataRequest, accountId int) (*requests.UserDataResponse, error)
	}

	IAddUserTestAnswersUseCase interface {
		Add(context context.Context, answers *requests.UserAnswersRequest, accountId int) (*requests.UserAnswersResponse, error)
	}

	ICheckIfUserHasPassedTestYetUseCase interface {
		Check(context context.Context, accountId int) (bool, error)
	}

	IAddUserXpChangeUseCase interface {
		Add(context context.Context, accountId int, request *requests.AddSkillChangeRequest) (*requests.UserDataResponse, error)
	}
)
