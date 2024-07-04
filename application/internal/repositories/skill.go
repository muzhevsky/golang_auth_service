package repositories

import (
	"context"
	"smartri_app/internal/entities/user_data"
	"smartri_app/internal/infrastructure/datasources"
)

type skillRepository struct {
	selectAllSkillsCommand                   datasources.ISelectAllSkillsCommand
	selectSkillsByAccountIdCommand           datasources.ISelectSkillsByAccountIdCommand
	selectSkillNormalizationBySkillIdCommand datasources.ISelectSkillNormalizationBySkillIdCommand
}

func NewSkillRepository(
	skillDataSource datasources.ISelectAllSkillsCommand,
	selectSkillsByAccountIdCommand datasources.ISelectSkillsByAccountIdCommand,
	skillNormalizationDataSource datasources.ISelectSkillNormalizationBySkillIdCommand) *skillRepository {
	return &skillRepository{
		selectAllSkillsCommand:                   skillDataSource,
		selectSkillsByAccountIdCommand:           selectSkillsByAccountIdCommand,
		selectSkillNormalizationBySkillIdCommand: skillNormalizationDataSource}
}

func (s *skillRepository) GetAllSkills(context context.Context) ([]*user_data.Skill, error) {
	return s.selectAllSkillsCommand.Execute(context)
}

func (s *skillRepository) GetSkillNormalizationBySkillId(context context.Context, skillId int) (*user_data.SkillNormalization, error) {
	return s.selectSkillNormalizationBySkillIdCommand.Execute(context, skillId)
}

func (s *skillRepository) GetSkillsByAccountId(ctx context.Context, accountId int) (*user_data.UserSkills, error) {
	return s.selectSkillsByAccountIdCommand.Execute(ctx, accountId)
}
