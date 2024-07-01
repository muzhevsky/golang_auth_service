package datasources

import (
	"context"
	"smartri_app/internal/entities"
)

type (
	ISelectAllQuestionsCommand interface {
		Execute(context context.Context) ([]*entities.Question, error)
	}

	ISelectQuestionByIdCommand interface {
		Execute(context context.Context, id int) (*entities.Question, error)
	}

	ISelectAnswerValuesByAnswerIdCommand interface {
		Execute(context context.Context, answerId int) ([]*entities.AnswerValue, error)
	}

	ISelectAnswerByIdCommand interface {
		Execute(context context.Context, id int) (*entities.Answer, error)
	}

	ISelectAnswersByQuestionIdCommand interface {
		Execute(context context.Context, questionId int) ([]*entities.Answer, error)
	}

	ISelectUserDataByAccountIdCommand interface {
		Execute(context context.Context, id int) (*entities.UserData, error)
	}

	IInsertUserDataCommand interface {
		Execute(context context.Context, user *entities.UserData) error
	}

	IUpdateUserDataByAccountIdCommand interface {
		Execute(context context.Context, newUser *entities.UserData) (*entities.UserData, error)
	}

	ISelectSkillsByAccountIdCommand interface {
		Execute(context context.Context, accountId int) ([]*entities.UserSkills, error)
	}

	IInsertUserTestResultsCommand interface {
		Execute(
			context context.Context,
			answers *entities.UserTestAnswers,
			changes []*entities.SkillChange,
			userSkills []*entities.UserSkills,
			userData *entities.UserData) error
	}

	ISelectAllSkillsCommand interface {
		Execute(context context.Context) ([]*entities.Skill, error)
	}

	ISelectAllSkillsNormalizationCommand interface {
		Execute(context context.Context) ([]*entities.SkillNormalization, error)
	}

	ISelectSkillNormalizationBySkillIdCommand interface {
		Execute(context context.Context, skillId int) (*entities.SkillNormalization, error)
	}

	ISelectSkillChangesByAccountIdCommand interface {
		Execute(context context.Context, id int) ([]*entities.SkillChange, error)
	}

	ISelectSkillChangesByAccountIdAndActionIdCommand interface {
		Execute(context context.Context, accountId int, actionId int) ([]*entities.SkillChange, error)
	}

	ISelectUserAnswersByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (*entities.UserTestAnswers, error)
	}

	ICheckIfUserHasAnswersByAccountIdCommand interface {
		Execute(context context.Context, accountId int) (bool, error)
	}

	IApplySkillChangesByAccountIdCommand interface {
		Execute(context context.Context, skills *entities.UserSkills, userData *entities.UserData, change *entities.SkillChange) error
	}
)
