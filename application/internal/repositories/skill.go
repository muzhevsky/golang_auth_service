package repositories

import (
	"context"
	"smartri_app/internal/entities"
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

func (s *skillRepository) GetAllSkills(context context.Context) ([]*entities.Skill, error) {
	return s.selectAllSkillsCommand.Execute(context)
}

func (s *skillRepository) GetSkillNormalizationBySkillId(context context.Context, skillId int) (*entities.SkillNormalization, error) {
	return s.selectSkillNormalizationBySkillIdCommand.Execute(context, skillId)
}

func (s *skillRepository) GetSkillsByAccountId(ctx context.Context, accountId int) ([]*entities.UserSkills, error) {
	return s.selectSkillsByAccountIdCommand.Execute(ctx, accountId)
}
