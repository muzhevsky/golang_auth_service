package answers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type selectAnswerValuesByAnswerIdCommand struct {
	client *postgres.Client
}

func NewSelectAnswerValuesByAnswerIdCommand(client *postgres.Client) datasources.ISelectAnswerValuesByAnswerIdCommand {
	return &selectAnswerValuesByAnswerIdCommand{client: client}
}

func (a *selectAnswerValuesByAnswerIdCommand) Execute(context context.Context, answerId int) ([]*entities.AnswerValue, error) {
	sql, args, err := query_builders.NewSelectAnswerValuesByAnswerIdQuery(&a.client.Builder, answerId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entities.AnswerValue{}, nil
		}
		return nil, nil
	}

	rows, err := a.client.Pool.Query(context, sql, args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	result := make([]*entities.AnswerValue, 0)

	for rows.Next() {
		value := entities.AnswerValue{AnswerId: answerId}
		err = rows.Scan(&value.Id, &value.SkillId, &value.Points)
		if err != nil {
			return nil, err
		}
		result = append(result, &value)
	}

	return result, err
}
