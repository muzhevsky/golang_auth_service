package skill_changes

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/skills"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectSkillChangesByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewSelectSkillChangesByAccountIdPGCommand(client *postgres.Client) *selectSkillChangesByAccountIdPGCommand {
	return &selectSkillChangesByAccountIdPGCommand{client: client}
}

func (s *selectSkillChangesByAccountIdPGCommand) Execute(context context.Context, accountId int) ([]*skills.SkillChange, error) {
	sql, args, err := query_builders.NewSelectSkillChangesByAccountIdQuery(&s.client.Builder, accountId)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*skills.SkillChange{}, nil
		}
		return nil, err
	}

	result := make([]*skills.SkillChange, 0)

	for rows.Next() {
		change := skills.SkillChange{AccountId: accountId}
		err = rows.Scan(&change.Id, &change.SkillId, change.Date, &change.ActionId, &change.Points)
		if err != nil {
			return nil, err
		}
		result = append(result, &change)
	}

	return result, nil
}
