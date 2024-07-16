package datasources

import (
	"context"
	"smartri_app/internal/entities/avatar_entities"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/entities/user_data_entities"
)

type (
	ISelectAllQuestionsCommand interface {
		Execute(context context.Context) ([]*test_entities.Question, error)
	}

	ISelectQuestionByIdCommand interface {
		Execute(context context.Context, id int) (*test_entities.Question, error)
	}

	ISelectAnswerValuesByAnswerIdCommand interface {
		Execute(context context.Context, answerId int) ([]*test_entities.AnswerValue, error)
	}

	ISelectAnswerByIdCommand interface {
		Execute(context context.Context, id int) (*test_entities.Answer, error)
	}

	ISelectAnswersByQuestionIdCommand interface {
		Execute(context context.Context, questionId int) ([]*test_entities.Answer, error)
	}

	ISelectUserDataByAccountIdCommand interface {
		Execute(context context.Context, id int) (*user_data_entities.UserData, error)
	}

	IInsertUserDataCommand interface {
		Execute(context context.Context, user *user_data_entities.UserData) error
	}

	IUpdateUserDataByAccountIdCommand interface {
		Execute(context context.Context, newUser *user_data_entities.UserData) (*user_data_entities.UserData, error)
	}

	ISelectSkillsByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (*skills_entities.UserSkills, error)
	}

	IInsertUserTestResultsCommand interface {
		Execute(
			context context.Context,
			answers *test_entities.UserTestAnswers,
			changes []*skills_entities.SkillChange,
			userSkills *skills_entities.UserSkills,
			userData *user_data_entities.UserData) error
	}

	ISelectAllSkillsCommand interface {
		Execute(context context.Context) ([]*skills_entities.Skill, error)
	}

	ISelectAllSkillsNormalizationCommand interface {
		Execute(context context.Context) ([]*skills_entities.SkillNormalization, error)
	}

	ISelectSkillNormalizationBySkillIdCommand interface {
		Execute(context context.Context, skillId int) (*skills_entities.SkillNormalization, error)
	}

	ISelectSkillChangesByAccountIdCommand interface {
		Execute(context context.Context, id int) ([]*skills_entities.SkillChange, error)
	}

	ISelectSkillChangesByAccountIdAndActionIdCommand interface {
		Execute(context context.Context, accountId int, actionId int) ([]*skills_entities.SkillChange, error)
	}

	ISelectUserAnswersByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (*test_entities.UserTestAnswers, error)
	}

	ICheckIfUserHasAnswersByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (bool, error)
	}

	IApplySkillChangesByAccountIdCommand interface {
		Execute(context context.Context, skills *skills_entities.UserSkill, userData *user_data_entities.UserData, change *skills_entities.SkillChange) error
	}

	ISelectAvatarByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (*avatar_entities.Avatar, error)
	}

	IInsertAvatarCommand interface {
		Execute(context context.Context, avatar *avatar_entities.Avatar) error
	}

	IUpdateAvatarCommand interface {
		Execute(context context.Context, accountId int, avatar *avatar_entities.Avatar) (*avatar_entities.Avatar, error)
	}
)
