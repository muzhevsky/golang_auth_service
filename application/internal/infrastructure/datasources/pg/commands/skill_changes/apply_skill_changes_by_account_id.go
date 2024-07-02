package skill_changes

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type applySkillChangesByAccountId struct {
	client *postgres.Client
}

func NewApplySkillChangesByAccountId(client *postgres.Client) *applySkillChangesByAccountId {
	return &applySkillChangesByAccountId{client: client}
}

func (a *applySkillChangesByAccountId) Execute(context context.Context, skills *entities.UserSkills, userData *entities.UserData, change *entities.SkillChange) error {
	sql, args, err := query_builders.NewUpdateUserSkillsQuery(&a.client.Builder, skills)
	if err != nil {
		return err
	}

	tx, err := a.client.Pool.Begin(context)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	sql, args, err = query_builders.NewInsertSkillChangesQuery(&a.client.Builder, change)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	sql, args, err = query_builders.NewUpdateUserDataByAccountIdQuery(&a.client.Builder, userData)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	err = tx.Commit(context)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	return nil
}
