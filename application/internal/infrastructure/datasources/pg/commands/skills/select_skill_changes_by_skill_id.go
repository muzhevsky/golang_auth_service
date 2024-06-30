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

type selectSkillChangesByAccountIdAndActionId struct {
	client *postgres.Client
}

func NewSelectSkillChangesByAccountIdAndActionId(client *postgres.Client) datasources.ISelectSkillChangesByAccountIdAndActionIdCommand {
	return &selectSkillChangesByAccountIdAndActionId{client: client}
}

func (s *selectSkillChangesByAccountIdAndActionId) Execute(context context.Context, accountId int, actionId int) ([]*entities.SkillChange, error) {
	sql, args, err := query_builders.NewSelectSkillChangesByAccountIdAndActionIdQuery(&s.client.Builder, accountId, actionId)

	if err != nil {
		return nil, err
	}

	rows, err := s.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entities.SkillChange{}, nil
		}
		return nil, err
	}

	result := make([]*entities.SkillChange, 0)

	for rows.Next() {
		change := entities.SkillChange{AccountId: accountId, ActionId: actionId}
		err = rows.Scan(&change.Id, &change.SkillId, change.Date, &change.Points)
		if err != nil {
			return nil, err
		}
		result = append(result, &change)
	}

	return result, nil
}
