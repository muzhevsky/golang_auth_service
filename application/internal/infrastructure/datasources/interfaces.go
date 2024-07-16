package datasources

import (
	"context"
	"smartri_app/internal/entities/avatar"
	"smartri_app/internal/entities/skills"
	"smartri_app/internal/entities/test"
	"smartri_app/internal/entities/user_data"
)

type (
	ISelectAllQuestionsCommand interface {
		Execute(context context.Context) ([]*test.Question, error)
	}

	ISelectQuestionByIdCommand interface {
		Execute(context context.Context, id int) (*test.Question, error)
	}

	ISelectAnswerValuesByAnswerIdCommand interface {
		Execute(context context.Context, answerId int) ([]*test.AnswerValue, error)
	}

	ISelectAnswerByIdCommand interface {
		Execute(context context.Context, id int) (*test.Answer, error)
	}

	ISelectAnswersByQuestionIdCommand interface {
		Execute(context context.Context, questionId int) ([]*test.Answer, error)
	}

	ISelectUserDataByAccountIdCommand interface {
		Execute(context context.Context, id int) (*user_data.UserData, error)
	}

	IInsertUserDataCommand interface {
		Execute(context context.Context, user *user_data.UserData) error
	}

	IUpdateUserDataByAccountIdCommand interface {
		Execute(context context.Context, newUser *user_data.UserData) (*user_data.UserData, error)
	}

	ISelectSkillsByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (*skills.UserSkills, error)
	}

	IInsertUserTestResultsCommand interface {
		Execute(
			context context.Context,
			answers *test.UserTestAnswers,
			changes []*skills.SkillChange,
			userSkills *skills.UserSkills,
			userData *user_data.UserData) error
	}

	ISelectAllSkillsCommand interface {
		Execute(context context.Context) ([]*skills.Skill, error)
	}

	ISelectAllSkillsNormalizationCommand interface {
		Execute(context context.Context) ([]*skills.SkillNormalization, error)
	}

	ISelectSkillNormalizationBySkillIdCommand interface {
		Execute(context context.Context, skillId int) (*skills.SkillNormalization, error)
	}

	ISelectSkillChangesByAccountIdCommand interface {
		Execute(context context.Context, id int) ([]*skills.SkillChange, error)
	}

	ISelectSkillChangesByAccountIdAndActionIdCommand interface {
		Execute(context context.Context, accountId int, actionId int) ([]*skills.SkillChange, error)
	}

	ISelectUserAnswersByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (*test.UserTestAnswers, error)
	}

	ICheckIfUserHasAnswersByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (bool, error)
	}

	IApplySkillChangesByAccountIdCommand interface {
		Execute(context context.Context, skills *skills.UserSkill, userData *user_data.UserData, change *skills.SkillChange) error
	}

	ISelectAvatarByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (*avatar.Avatar, error)
	}

	IInsertAvatarCommand interface {
		Execute(context context.Context, avatar *avatar.Avatar) error
	}

	IUpdateAvatarCommand interface {
		Execute(context context.Context, accountId int, avatar *avatar.Avatar) (*avatar.Avatar, error)
	}
)
