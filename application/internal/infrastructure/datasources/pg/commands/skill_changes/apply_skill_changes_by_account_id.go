package skill_changes

import (
	"context"
	"smartri_app/internal/entities/skills"
	"smartri_app/internal/entities/user_data"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type applySkillChangesByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewApplySkillChangesByAccountIdPGCommand(client *postgres.Client) *applySkillChangesByAccountIdPGCommand {
	return &applySkillChangesByAccountIdPGCommand{client: client}
}

func (a *applySkillChangesByAccountIdPGCommand) Execute(context context.Context, skills *skills.UserSkill, userData *user_data.UserData, change *skills.SkillChange) error {
	sql, args, err := query_builders.NewUpdateUserSkillsQuery(&a.client.Builder, userData.AccountId, skills)
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
