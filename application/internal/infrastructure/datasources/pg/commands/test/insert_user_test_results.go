package test

import (
	"context"
	"smartri_app/internal/entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type insertUserTestResultsCommand struct {
	client *postgres.Client
}

func NewInsertUserTestResultsCommand(
	client *postgres.Client) datasources.IInsertUserTestResultsCommand {
	return &insertUserTestResultsCommand{client: client}
}

func (u *insertUserTestResultsCommand) Execute(
	context context.Context,
	answers *entities.UserTestAnswers,
	changes []*entities.SkillChange,
	userSkills []*entities.UserSkills,
	userData *entities.UserData) error {
	insertTestResultsSQL, args, err := query_builders.NewInsertUserTestResultsQuery(&u.client.Builder, answers)
	if err != nil {
		return err
	}

	tx, err := u.client.Pool.Begin(context)

	if err != nil {
		tx.Rollback(context)
		return err
	}

	_, err = tx.Exec(context, insertTestResultsSQL, args...)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	for _, skill := range changes {
		insertSkillChangesSQL, args, err := query_builders.NewInsertSkillChangesQuery(&u.client.Builder, skill)
		if err != nil {
			tx.Rollback(context)
			return err
		}

		_, err = tx.Exec(context, insertSkillChangesSQL, args...)
		if err != nil {
			tx.Rollback(context)
			return err
		}
	}

	updateUserXpSQL, args, err := query_builders.NewUpdateUserDataByAccountIdQuery(&u.client.Builder, userData)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	_, err = tx.Exec(context, updateUserXpSQL, args...)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	for i := range userSkills {
		updateUserSkillsSQL, args, err := query_builders.NewUpdateUserSkillsByAccountIdQuery(&u.client.Builder, userSkills[i])
		if err != nil {
			tx.Rollback(context)
			return err
		}

		_, err = tx.Exec(context, updateUserSkillsSQL, args...)
	}

	err = tx.Commit(context)
	if err != nil {
		tx.Rollback(context)
		return err
	}

	return nil
}
