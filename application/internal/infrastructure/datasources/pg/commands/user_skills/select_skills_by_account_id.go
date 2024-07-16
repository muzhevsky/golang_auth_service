package user_skills

import (
	"context"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectSkillsByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectSkillsByAccountIdPGCommand(client *postgres.Client) datasources.ISelectSkillsByAccountIdCommand {
	return &selectSkillsByAccountIdPGCommand{client: client}
}

func (c *selectSkillsByAccountIdPGCommand) Execute(context context.Context, accountId int) (*skills_entities.UserSkills, error) {
	sql, args, err := query_builders.NewSelectUserSkillsByAccountIdQuery(&c.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	rows, err := c.client.Pool.Query(context, sql, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := skills_entities.UserSkills{AccountId: accountId}
	skills := make([]*skills_entities.UserSkill, 0)
	for rows.Next() {
		newUserSkill := skills_entities.UserSkill{}
		err = rows.Scan(&newUserSkill.SkillId, &newUserSkill.Xp)
		if err != nil {
			return nil, err
		}
		skills = append(skills, &newUserSkill)
	}
	result.Skills = skills
	return &result, nil
}
