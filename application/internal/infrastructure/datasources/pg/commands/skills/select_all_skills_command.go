package skills

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/skills"
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

func (s *selectAllSkillsPGCommand) Execute(context context.Context) ([]*skills.Skill, error) {
	sql, i, err := query_builders.NewSelectAllSkillsQuery(&s.client.Builder)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, i...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*skills.Skill{}, nil
		}
		return nil, err
	}
	skills := make([]*skills.Skill, 0)

	for rows.Next() {
		newSkill := skills.Skill{}
		err = rows.Scan(&newSkill.Id, &newSkill.Title)
		if err != nil {
			return nil, err

		}
		skills = append(skills, &newSkill)
	}

	return skills, nil
}
