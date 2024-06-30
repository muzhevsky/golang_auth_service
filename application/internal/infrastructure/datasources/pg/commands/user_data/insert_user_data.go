package user_data

import (
	"context"
	"fmt"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type insertUserDataCommand struct {
	client                 *postgres.Client
	selectAllSkillsCommand datasources.ISelectAllSkillsCommand
}

func NewInsertUserDataCommand(client *postgres.Client, selectAllSkillsCommand datasources.ISelectAllSkillsCommand) datasources.IInsertUserDataCommand {
	return &insertUserDataCommand{client: client, selectAllSkillsCommand: selectAllSkillsCommand}
}

func (u *insertUserDataCommand) Execute(context context.Context, user *entities.UserData) error {
	sql, args, err := query_builders.NewInsertUserDataQuery(&u.client.Builder, user)
	if err != nil {
		fmt.Println("1")
		return err
	}

	tx, err := u.client.Pool.Begin(context)

	if err != nil {
		fmt.Println("2")
		return err
	}

	_, err = tx.Exec(context, sql, args...)
	if err != nil {
		fmt.Println("3")
		tx.Rollback(context)
		return err
	}
	fmt.Println("ok")

	skills, err := u.selectAllSkillsCommand.Execute(context)

	for _, skill := range skills {
		sql, args, err = query_builders.NewInsertUserSkillsQuery(&u.client.Builder, &entities.UserSkills{
			AccountId: user.AccountId,
			SkillId:   skill.Id,
			Xp:        0,
		})

		if err != nil {
			fmt.Println("4")
			tx.Rollback(context)
			return err
		}

		_, err = tx.Exec(context, sql, args...)
		if err != nil {
			fmt.Println("5")
			tx.Rollback(context)
			return err
		}
	}
	tx.Commit(context)

	return nil
}
