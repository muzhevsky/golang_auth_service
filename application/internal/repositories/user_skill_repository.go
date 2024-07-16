package repositories

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/infrastructure/datasources"
)

type userSkillRepository struct {
	applySkillChangesByAccountIdCommand datasources.IApplySkillChangesByAccountIdCommand
}

func NewUserSkillRepository(applySkillChangesByAccountIdCommand datasources.IApplySkillChangesByAccountIdCommand) internal.IUserSkillsRepository {
	return &userSkillRepository{applySkillChangesByAccountIdCommand: applySkillChangesByAccountIdCommand}
}

func (u *userSkillRepository) ApplySkillChangesByAccountId(context context.Context, userSkills *skills_entities.UserSkill, userData *user_data_entities.UserData, change *skills_entities.SkillChange) error {
	return u.applySkillChangesByAccountIdCommand.Execute(context, userSkills, userData, change)
}
