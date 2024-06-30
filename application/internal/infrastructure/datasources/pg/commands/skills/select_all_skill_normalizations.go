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

type selectAllSkillNormalizationsCommand struct {
	client *postgres.Client
}

func NewSelectAllSkillNormalizationsCommand(client *postgres.Client) datasources.ISelectAllSkillsNormalizationCommand {
	return &selectAllSkillNormalizationsCommand{client: client}
}

func (s *selectAllSkillNormalizationsCommand) Execute(context context.Context) ([]*entities.SkillNormalization, error) {
	sql, i, err := query_builders.NewSelectAllSkillsNormalizationsQuery(&s.client.Builder)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, i...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entities.SkillNormalization{}, nil
		}
		return nil, err
	}
	skills := make([]*entities.SkillNormalization, 0)

	for rows.Next() {
		newSkill := entities.SkillNormalization{}
		err = rows.Scan(&newSkill.SkillId, &newSkill.Min, &newSkill.Max)
		if err != nil {
			return nil, err

		}
		skills = append(skills, &newSkill)
	}

	return skills, nil
}
