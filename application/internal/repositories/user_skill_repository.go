package repositories

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/entities/user_data"
	"smartri_app/internal/infrastructure/datasources"
)

type userSkillRepository struct {
	applySkillChangesByAccountIdCommand datasources.IApplySkillChangesByAccountIdCommand
}

func NewUserSkillRepository(applySkillChangesByAccountIdCommand datasources.IApplySkillChangesByAccountIdCommand) internal.IUserSkillsRepository {
	return &userSkillRepository{applySkillChangesByAccountIdCommand: applySkillChangesByAccountIdCommand}
}

func (u *userSkillRepository) ApplySkillChangesByAccountId(context context.Context, userSkills *user_data.UserSkill, userData *user_data.UserData, change *user_data.SkillChange) error {
	return u.applySkillChangesByAccountIdCommand.Execute(context, userSkills, userData, change)
}
