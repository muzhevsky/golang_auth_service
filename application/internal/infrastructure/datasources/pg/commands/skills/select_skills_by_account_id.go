package skills

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectSkillsByAccountIdCommand struct {
	client *postgres.Client
}

func NewSelectSkillsByAccountIdCommand(client *postgres.Client) *selectSkillsByAccountIdCommand {
	return &selectSkillsByAccountIdCommand{client: client}
}

func (c *selectSkillsByAccountIdCommand) Execute(context context.Context, accountId int) ([]*entities.UserSkills, error) {
	sql, args, err := query_builders.NewSelectSkillsByAccountIdQuery(&c.client.Builder, accountId)
	if err != nil {
		return nil, err
	}

	rows, err := c.client.Pool.Query(context, sql, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := make([]*entities.UserSkills, 0)
	for rows.Next() {
		newUserSkill := entities.UserSkills{AccountId: accountId}
		err = rows.Scan(&newUserSkill.SkillId, &newUserSkill.Xp)
		if err != nil {
			return nil, err
		}
		result = append(result, &newUserSkill)
	}
	return result, nil
}
