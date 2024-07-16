package internal

import (
	"context"
	requests2 "smartri_app/controllers/requests"
	"smartri_app/internal/entities/avatar"
	"smartri_app/internal/entities/skills"
	"smartri_app/internal/entities/test"
	"smartri_app/internal/entities/user_data"
)

type (
	ITestRepository interface {
		GetAllQuestions(context context.Context) ([]*test.Question, error)
		GetAnswersForQuestion(context context.Context, question *test.Question) ([]*test.Answer, error)
		GetAnswerWithValues(context context.Context, answerId int) (*test.Answer, error)
		GetAllQuestionsWithAnswers(context context.Context) ([]*test.Question, error)
		AddUserAnswersWithSkillChanges(context context.Context, results *test.UserTestAnswers, skills []*skills.SkillChange, userSkills *skills.UserSkills, data *user_data.UserData) error
	}

	ISkillRepository interface {
		GetAllSkills(context context.Context) ([]*skills.Skill, error)
		GetSkillNormalizationBySkillId(context context.Context, skillId int) (*skills.SkillNormalization, error)
		GetSkillsByAccountId(context context.Context, accountId int) (*skills.UserSkills, error)
	}

	IUserDataRepository interface {
		GetByAccountId(context context.Context, accountId int) (*user_data.UserData, error)
		Create(context context.Context, userData *user_data.UserData) error
		Update(context context.Context, details *user_data.UserData) (*user_data.UserData, error)
	}

	IUserSkillsRepository interface {
		ApplySkillChangesByAccountId(context context.Context, userSkill *skills.UserSkill, userData *user_data.UserData, change *skills.SkillChange) error
	}

	IUserAnswersRepository interface {
		CheckUserHasAnswers(context context.Context, accountId int) (bool, error)
	}

	IAvatarRepository interface {
		GetByAccountId(context context.Context, accountId int) (*avatar.Avatar, error)
		Create(context context.Context, avatar *avatar.Avatar) error
		Update(context context.Context, accountId int, avatar *avatar.Avatar) (*avatar.Avatar, error)
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
)
