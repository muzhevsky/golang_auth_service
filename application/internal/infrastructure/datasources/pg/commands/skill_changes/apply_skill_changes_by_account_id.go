package skill_changes

import (
	"context"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type applySkillChangesByAccountIdPGCommand struct {
	client *postgres.Client
}

func NewApplySkillChangesByAccountIdPGCommand(client *postgres.Client) datasources.IApplySkillChangesByAccountIdCommand {
	return &applySkillChangesByAccountIdPGCommand{client: client}
}

func (a *applySkillChangesByAccountIdPGCommand) Execute(context context.Context, skills *skills_entities.UserSkill, userData *user_data_entities.UserData, change *skills_entities.SkillChange) error {
	sql, args, err := query_builders.NewUpdateUserSkillsQuery(&a.client.Builder, userData.AccountId, skills)
	if err != nil {
		return err
	}

	tx, err := a.client.Pool.Begin(context)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	sql, args, err = query_builders.NewInsertSkillChangesQuery(&a.client.Builder, change)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	sql, args, err = query_builders.NewUpdateUserDataByAccountIdQuery(&a.client.Builder, userData)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	err = tx.Commit(context)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	return nil
}
