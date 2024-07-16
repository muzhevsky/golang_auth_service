package answers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAnswerValuesByAnswerIdPGCommand struct {
	client *postgres.Client
}

func NewSelectAnswerValuesByAnswerIdPGCommand(client *postgres.Client) datasources.ISelectAnswerValuesByAnswerIdCommand {
	return &selectAnswerValuesByAnswerIdPGCommand{client: client}
}

func (a *selectAnswerValuesByAnswerIdPGCommand) Execute(context context.Context, answerId int) ([]*test_entities.AnswerValue, error) {
	sql, args, err := query_builders.NewSelectAnswerValuesByAnswerIdQuery(&a.client.Builder, answerId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*test_entities.AnswerValue{}, nil
		}
		return nil, nil
	}

	rows, err := a.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*test_entities.AnswerValue, 0)

	for rows.Next() {
		value := test_entities.AnswerValue{}
		err = rows.Scan(&value.Id, &value.AnswerId, &value.SkillId, &value.Points)
		if err != nil {
			return nil, err
		}
		result = append(result, &value)
	}

	return result, err
}
