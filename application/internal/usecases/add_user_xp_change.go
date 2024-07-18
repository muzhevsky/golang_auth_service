package usecases

import (
	"context"
	"fmt"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/errs"
	"time"
)

type addUserXpChange struct {
	skillRepo      internal.ISkillRepository
	userDataRepo   internal.IUserDataRepository
	userSkillsRepo internal.IUserSkillsRepository
}

func NewAddUserXpChange(
	skillRepository internal.ISkillRepository,
	userDataRepository internal.IUserDataRepository,
	userSkillsRepository internal.IUserSkillsRepository) internal.IAddUserXpChangeUseCase {
	return &addUserXpChange{
		skillRepo:      skillRepository,
		userDataRepo:   userDataRepository,
		userSkillsRepo: userSkillsRepository,
	}
}

func (uc *addUserXpChange) Add(context context.Context, accountId int, request *requests.AddSkillChangeRequest) (*requests.UserDataResponse, error) {
	skillsByAccount, err := uc.skillRepo.GetSkillsByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	userData, err := uc.userDataRepo.GetByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	var skill *skills_entities.UserSkill
	for i := range skillsByAccount.Skills {
		s := skillsByAccount.Skills[i]
		if s.SkillId == request.SkillId {
			prev := s.Xp
			s.Xp += request.Points
			if s.Xp > maxPoints {
				request.Points = maxPoints - prev
				s.Xp = maxPoints
			}
			userData.XP += user_data_entities.XP(request.Points)
			skill = s
			break
		}
	}

	if skill == nil {
		return nil, fmt.Errorf("%w skill not found", errs.EntityNotFoundError)
	}

	err = uc.userSkillsRepo.ApplySkillChangesByAccountId(context, skill, userData, &skills_entities.SkillChange{
		AccountId: accountId,
		SkillId:   skill.SkillId,
		ActionId:  1, // todo а надо ли?
		Date:      time.Time{},
		Points:    request.Points,
	})
	if err != nil {
		return nil, err
	}

	return &requests.UserDataResponse{
		Age:    int(userData.Age),
		Gender: string(userData.Gender),
		XP:     int(userData.XP),
	}, nil
}
