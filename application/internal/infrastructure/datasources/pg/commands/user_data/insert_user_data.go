package user_data

import (
	"context"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type insertUserDataPGCommand struct {
	client                 *postgres.Client
	selectAllSkillsCommand datasources.ISelectAllSkillsCommand
}

func NewInsertUserDataPGCommand(client *postgres.Client, selectAllSkillsCommand datasources.ISelectAllSkillsCommand) datasources.IInsertUserDataCommand {
	return &insertUserDataPGCommand{client: client, selectAllSkillsCommand: selectAllSkillsCommand}
}

func (u *insertUserDataPGCommand) Execute(context context.Context, userData *user_data_entities.UserData) error {
	sql, args, err := query_builders.NewInsertUserDataQuery(&u.client.Builder, userData)
	if err != nil {
		return err
	}

	tx, err := u.client.Pool.Begin(context)

	if err != nil {
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	skills, err := u.selectAllSkillsCommand.Execute(context)

	for _, skill := range skills {
		sql, args, err = query_builders.NewInsertUserSkillsQuery(&u.client.Builder, userData.AccountId, &skills_entities.UserSkill{
			SkillId: skill.Id,
			Xp:      0,
		})

		if err != nil {
			tx.Rollback(context)
			return err
		}

		_, err = tx.Exec(context, sql, args...)
		if err != nil {
			tx.Rollback(context)
			return err
		}
	}
	tx.Commit(context)

	return nil
}
