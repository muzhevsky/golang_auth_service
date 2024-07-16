package skills

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAllSkillsPGCommand struct {
	client *postgres.Client
}

func NewSelectAllSkillsPGCommand(client *postgres.Client) datasources.ISelectAllSkillsCommand {
	return &selectAllSkillsPGCommand{client: client}
}

func (s *selectAllSkillsPGCommand) Execute(context context.Context) ([]*skills_entities.Skill, error) {
	sql, i, err := query_builders.NewSelectAllSkillsQuery(&s.client.Builder)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, i...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*skills_entities.Skill{}, nil
		}
		return nil, err
	}
	skills := make([]*skills_entities.Skill, 0)

	for rows.Next() {
		newSkill := skills_entities.Skill{}
		err = rows.Scan(&newSkill.Id, &newSkill.Title)
		if err != nil {
			return nil, err

		}
		skills = append(skills, &newSkill)
	}

	return skills, nil
}
