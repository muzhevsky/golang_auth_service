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

type selectAllSkillNormalizationsBySkillIdCommand struct {
	client *postgres.Client
}

func NewSelectSkillNormalizationBySkillIdCommand(client *postgres.Client) datasources.ISelectSkillNormalizationBySkillIdCommand {
	return &selectAllSkillNormalizationsBySkillIdCommand{client: client}
}

func (s *selectAllSkillNormalizationsBySkillIdCommand) Execute(context context.Context, skillId int) (*entities.SkillNormalization, error) {
	sql, i, err := query_builders.NewSelectSkillNormalizationsBySkillIdQuery(&s.client.Builder, skillId)

	if err != nil {
		return nil, err
	}

	row := s.client.Pool.QueryRow(context, sql, i...)
	result := entities.SkillNormalization{}
	err = row.Scan(&result.SkillId, &result.Min, &result.Max)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
