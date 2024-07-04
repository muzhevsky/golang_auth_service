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

type selectAllSkillNormalizationsBySkillIdPGCommand struct {
	client *postgres.Client
}

func NewSelectSkillNormalizationBySkillIdPGCommand(client *postgres.Client) datasources.ISelectSkillNormalizationBySkillIdCommand {
	return &selectAllSkillNormalizationsBySkillIdPGCommand{client: client}
}

func (s *selectAllSkillNormalizationsBySkillIdPGCommand) Execute(context context.Context, skillId int) (*user_data.SkillNormalization, error) {
	sql, i, err := query_builders.NewSelectSkillNormalizationsBySkillIdQuery(&s.client.Builder, skillId)

	if err != nil {
		return nil, err
	}

	row := s.client.Pool.QueryRow(context, sql, i...)
	result := user_data.SkillNormalization{}
	err = row.Scan(&result.SkillId, &result.Min, &result.Max)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
