package internal

import (
	"context"
	requests2 "smartri_app/controllers/requests"
	"smartri_app/internal/entities/avatar_entities"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/entities/user_data_entities"
)

type (
	ITestRepository interface {
		GetAllQuestions(context context.Context) ([]*test_entities.Question, error)
		GetAnswersForQuestion(context context.Context, question *test_entities.Question) ([]*test_entities.Answer, error)
		GetAnswerWithValues(context context.Context, answerId int) (*test_entities.Answer, error)
		GetAllQuestionsWithAnswers(context context.Context) ([]*test_entities.Question, error)
		AddUserAnswersWithSkillChanges(context context.Context, results *test_entities.UserTestAnswers, skills []*skills_entities.SkillChange, userSkills *skills_entities.UserSkills, data *user_data_entities.UserData) error
	}

	ISkillRepository interface {
		GetAllSkills(context context.Context) ([]*skills_entities.Skill, error)
		GetSkillNormalizationBySkillId(context context.Context, skillId int) (*skills_entities.SkillNormalization, error)
		GetSkillsByAccountId(context context.Context, accountId int) (*skills_entities.UserSkills, error)
	}

	IUserDataRepository interface {
		GetByAccountId(context context.Context, accountId int) (*user_data_entities.UserData, error)
		Create(context context.Context, userData *user_data_entities.UserData) error
		Update(context context.Context, details *user_data_entities.UserData) (*user_data_entities.UserData, error)
	}

	IUserSkillsRepository interface {
		ApplySkillChangesByAccountId(context context.Context, userSkill *skills_entities.UserSkill, userData *user_data_entities.UserData, change *skills_entities.SkillChange) error
	}

	IUserAnswersRepository interface {
		CheckUserHasAnswers(context context.Context, accountId int) (bool, error)
	}

	IAvatarRepository interface {
		GetByAccountId(context context.Context, accountId int) (*avatar_entities.Avatar, error)
		Create(context context.Context, avatar *avatar_entities.Avatar) error
		Update(context context.Context, accountId int, avatar *avatar_entities.Avatar) (*avatar_entities.Avatar, error)
	}

	IInitOrUpdateUserDataUseCase interface {
		InitOrUpdate(context context.Context, data *requests2.UserDataRequest, accountId int) (*requests2.UserDataResponse, error)
	}

	IAddUserTestAnswersUseCase interface {
		Add(context context.Context, answers *requests2.UserAnswersRequest, accountId int) (*requests2.UserAnswersResponse, error)
	}

	ICheckIfUserHasPassedTestYetUseCase interface {
		Check(context context.Context, accountId int) (bool, error)
	}

	IAddUserXpChangeUseCase interface {
		Add(context context.Context, accountId int, request *requests2.AddSkillChangeRequest) (*requests2.UserDataResponse, error)
	}

	IInitOrUpdateAvatarUseCase interface {
		InitOrUpdate(context context.Context, accountId int, request *requests2.AvatarRequest) error
	}

	IGetUserAvatarUseCase interface {
		GetAvatar(context context.Context, accountId int) (*requests2.AvatarRequest, error)
	}

	IGetUserDataUseCase interface {
		GetUserData(context context.Context, accountId int) (*requests2.UserDataResponse, error)
	}
)
