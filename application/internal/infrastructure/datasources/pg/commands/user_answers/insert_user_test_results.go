package user_answers

import (
	"context"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/infrastructure/datasources"
	"smartri_app/internal/infrastructure/datasources/pg"
	"smartri_app/internal/infrastructure/datasources/pg/query_builders"
	"smartri_app/pkg/postgres"
)

type insertUserTestResultsPGCommand struct {
	client *postgres.Client
}

func NewInsertUserTestResultsPGCommand(
	client *postgres.Client) datasources.IInsertUserTestResultsCommand {
	return &insertUserTestResultsPGCommand{client: client}
}

func (u *insertUserTestResultsPGCommand) Execute(
	context context.Context,
	answers *test_entities.UserTestAnswers,
	changes []*skills_entities.SkillChange,
	userSkills *skills_entities.UserSkills,
	userData *user_data_entities.UserData) error {
	insertTestResultsSQL, args, err := query_builders.NewInsertUserTestResultsQuery(&u.client.Builder, answers)
	if err != nil {
		return err
	}

	tx, err := u.client.Pool.Begin(context)

	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	_, err = tx.Exec(context, insertTestResultsSQL, args...)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	for _, skill := range changes {
		insertSkillChangesSQL, args, err := query_builders.NewInsertSkillChangesQuery(&u.client.Builder, skill)
		if pg.WrapError(context, err, tx) != nil {
			return err
		}

		_, err = tx.Exec(context, insertSkillChangesSQL, args...)
		if pg.WrapError(context, err, tx) != nil {
			return err
		}
	}

	updateUserXpSQL, args, err := query_builders.NewUpdateUserDataByAccountIdQuery(&u.client.Builder, userData)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	_, err = tx.Exec(context, updateUserXpSQL, args...)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	for i := range userSkills.Skills {
		updateUserSkillsSQL, args, err := query_builders.NewUpdateUserSkillsByAccountIdQuery(&u.client.Builder, userData.AccountId, userSkills.Skills[i])
		if pg.WrapError(context, err, tx) != nil {
			return err
		}

		_, err = tx.Exec(context, updateUserSkillsSQL, args...)
	}

	err = tx.Commit(context)
	if pg.WrapError(context, err, tx) != nil {
		return err
	}

	return nil
}
