package skills

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAllSkillsCommand struct {
	client *postgres.Client
}

func NewSelectAllSkillsCommand(client *postgres.Client) datasources.ISelectAllSkillsCommand {
	return &selectAllSkillsCommand{client: client}
}

func (s *selectAllSkillsCommand) Execute(context context.Context) ([]*entities.Skill, error) {
	sql, i, err := query_builders.NewSelectAllSkillsQuery(&s.client.Builder)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, i...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entities.Skill{}, nil
		}
		return nil, err
	}
	skills := make([]*entities.Skill, 0)

	for rows.Next() {
		newSkill := entities.Skill{}
		err = rows.Scan(&newSkill.Id, &newSkill.Title)
		if err != nil {
			return nil, err

		}
		skills = append(skills, &newSkill)
	}

	return skills, nil
}
