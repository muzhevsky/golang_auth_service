package skills

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/user_data"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAllSkillNormalizationsPGCommand struct {
	client *postgres.Client
}

func NewSelectAllSkillNormalizationsPGCommand(client *postgres.Client) datasources.ISelectAllSkillsNormalizationCommand {
	return &selectAllSkillNormalizationsPGCommand{client: client}
}

func (s *selectAllSkillNormalizationsPGCommand) Execute(context context.Context) ([]*user_data.SkillNormalization, error) {
	sql, i, err := query_builders.NewSelectAllSkillsNormalizationsQuery(&s.client.Builder)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, i...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*user_data.SkillNormalization{}, nil
		}
		return nil, err
	}
	skills := make([]*user_data.SkillNormalization, 0)

	for rows.Next() {
		newSkill := user_data.SkillNormalization{}
		err = rows.Scan(&newSkill.SkillId, &newSkill.Min, &newSkill.Max)
		if err != nil {
			return nil, err

		}
		skills = append(skills, &newSkill)
	}

	return skills, nil
}
