package repositories

import (
	"authorization/internal"
	"authorization/internal/entities"
	"authorization/internal/infrastructure/datasources"
	"context"
)

type verificationRepo struct {
	ds datasources.IVerificationDataSource
}

func NewVerificationRepo(ds datasources.IVerificationDataSource) internal.IVerificationRepository {
	return &verificationRepo{ds}
}

func (repo *verificationRepo) Create(context context.Context, verification *entities.Verification) (int, error) {
	return repo.ds.Create(context, verification)
}

func (repo *verificationRepo) FindById(context context.Context, id int) (*entities.Verification, error) {
	return repo.ds.SelectById(context, id)
}

func (repo *verificationRepo) FindByAccountId(context context.Context, userId int) ([]*entities.Verification, error) {
	return repo.ds.SelectByUserId(context, userId)
}

func (repo *verificationRepo) Clear(context context.Context, userId int) error {
	verifications, err := repo.ds.SelectByUserId(context, userId)
	if err != nil {
		return err
	}
	for _, verification := range verifications {
		err = repo.ds.DeleteById(context, verification.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
