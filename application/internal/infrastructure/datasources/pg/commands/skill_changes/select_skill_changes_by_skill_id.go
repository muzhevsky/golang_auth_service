package skill_changes

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectSkillChangesByAccountIdAndActionIdPGCommand struct {
	client *postgres.Client
}

func NewSelectSkillChangesByAccountIdAndActionIdPGCommand(client *postgres.Client) datasources.ISelectSkillChangesByAccountIdAndActionIdCommand {
	return &selectSkillChangesByAccountIdAndActionIdPGCommand{client: client}
}

func (s *selectSkillChangesByAccountIdAndActionIdPGCommand) Execute(context context.Context, accountId int, actionId int) ([]*skills_entities.SkillChange, error) {
	sql, args, err := query_builders.NewSelectSkillChangesByAccountIdAndActionIdQuery(&s.client.Builder, accountId, actionId)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*skills_entities.SkillChange{}, nil
		}
		return nil, err
	}

	result := make([]*skills_entities.SkillChange, 0)

	for rows.Next() {
		change := skills_entities.SkillChange{}
		err = rows.Scan(&change.Id, &change.AccountId, &change.SkillId, &change.Date, &change.ActionId, &change.Points)
		if err != nil {
			return nil, err
		}
		result = append(result, &change)
	}

	return result, nil
}
