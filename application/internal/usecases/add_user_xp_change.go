package usecases

import (
	"context"
	_ "smartri_app/docs"
	"smartri_app/internal"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/entities"
	"time"
)

type addUserXpChange struct {
	skillRepository internal.ISkillRepository
	userRepository  internal.IUserDataRepository
}

func NewAddUserXpChange(
	skillRepository internal.ISkillRepository,
	userRepository internal.IUserDataRepository) *addUserXpChange {
	return &addUserXpChange{
		skillRepository: skillRepository,
		userRepository:  userRepository,
	}
}

// Add godoc
// @Summary      добавляет опыт пользователю
// @Description  добавляет опыт пользователю по скиллам используя access token
// @Accept       json
// @Produce      json
// @Param request body requests.AddSkillChangeRequest true "request format"
// @Param Authorization header string true "access token"
// @Success      200  {object} requests.UserDataResponse
// @Failure 400 {object} middleware.ErrorResponse "некорректный формат запроса"
// @Failure 401 {object} middleware.ErrorResponse "ошибка аутентификации"
// @Failure 500 {object} middleware.ErrorResponse "внутренняя ошибка сервера"
// @Router       /user/xp [post]
func (uc *addUserXpChange) Add(context context.Context, accountId int, request *requests.AddSkillChangeRequest) (*requests.UserDataResponse, error) {
	skills, err := uc.skillRepository.GetSkillsByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	userData, err := uc.userRepository.GetDataByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	var skill *entities.UserSkills
	for i := range skills {
		if skills[i].SkillId == request.SkillId {
			skills[i].Xp += request.Points
			if skills[i].Xp > maxPoints {
				request.Points = skills[i].Xp - maxPoints
				skills[i].Xp = maxPoints
			}
			userData.XP += request.Points
			skill = skills[i]
			break
		}
	}

	skill.AccountId = accountId
	err = uc.userRepository.ApplySkillChangesByAccountId(context, skill, userData, &entities.SkillChange{
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
		Age:    userData.Age,
		Gender: userData.Gender,
		XP:     userData.XP,
	}, nil
}
