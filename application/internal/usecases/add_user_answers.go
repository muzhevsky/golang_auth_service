package usecases

import (
	"context"
	"smartri_app/internal"
	"smartri_app/internal/controllers/requests"
	"smartri_app/internal/entities"
	"smartri_app/internal/errs"
	"time"
)

const maxPoints = 200

type addUserAnswers struct {
	testRepo  internal.ITestRepository
	skillRepo internal.ISkillRepository
	userRepo  internal.IUserDataRepository
}

func NewAddUserAnswers(
	testRepo internal.ITestRepository,
	skillRepo internal.ISkillRepository,
	userRepo internal.IUserDataRepository) internal.IAddUserTestAnswersUseCase {
	return &addUserAnswers{testRepo: testRepo, skillRepo: skillRepo, userRepo: userRepo}
}

func (a *addUserAnswers) Add(context context.Context, answers *requests.UserAnswersRequest, accountId int) (*requests.UserAnswersResponse, error) {
	userData, err := a.userRepo.GetUserDataByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, errs.UserDataNotFoundError
	}

	userHasAnswers, err := a.userRepo.CheckUserHasAnswers(context, accountId)
	if err != nil {
		return nil, err
	}
	if userHasAnswers {
		return nil, errs.UserHasAlreadyPassedTestError
	}

	skills, err := a.skillRepo.GetAllSkills(context)
	if err != nil {
		return nil, err
	}
	normalizationBySkillIdMap, err := a.getNormalizationMap(context, skills, accountId)
	if err != nil {
		return nil, err
	}

	entityUserAnswers := a.getEntityAnswers(accountId, answers)
	answersWithValues, err := a.getAnswersWithValues(context, entityUserAnswers)
	if err != nil {
		return nil, err
	}

	skillIdPointsMap := a.getSkillIdPointsMap(answersWithValues)
	userSkills, skillChanges := a.getUserSkillsAndSkillChanges(skillIdPointsMap, normalizationBySkillIdMap, accountId)
	userData.XP = a.getNewUserXpValue(userSkills)

	err = a.testRepo.AddUserTestResults(context, entityUserAnswers, skillChanges, userSkills, userData)
	if err != nil {
		return nil, err
	}

	result, err := a.skillRepo.GetSkillsByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	return &requests.UserAnswersResponse{
		AccountId: accountId,
		Skills:    result,
		TotalExp:  userData.XP,
	}, nil
}

func (a *addUserAnswers) getEntityAnswers(accountId int, answers *requests.UserAnswersRequest) *entities.UserTestAnswers {
	result := &entities.UserTestAnswers{accountId, make([]entities.UserTestAnswer, 0)}
	for i := range answers.Answers {
		result.Answers = append(result.Answers, entities.UserTestAnswer{
			QuestionId: answers.Answers[i].QuestionId,
			AnswerId:   answers.Answers[i].AnswerId,
		})
	}

	return result
}

func (a *addUserAnswers) getNormalizationMap(context context.Context, skills []*entities.Skill, accountId int) (map[int]*entities.SkillNormalization, error) {
	result := make(map[int]*entities.SkillNormalization)
	for i := range skills {
		r, err := a.skillRepo.GetSkillNormalizationBySkillId(context, skills[i].Id)
		result[i] = r
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (a *addUserAnswers) getAnswersWithValues(context context.Context, entityAnswers *entities.UserTestAnswers) ([]*entities.Answer, error) {
	result := make([]*entities.Answer, 0)
	for i := 0; i < len(entityAnswers.Answers); i++ {
		answer, err := a.testRepo.GetAnswerWithValues(context, entityAnswers.Answers[i].AnswerId)
		if err != nil {
			return nil, err
		}

		result = append(result, answer)
	}

	return result, nil
}

func (a *addUserAnswers) getSkillIdPointsMap(answersWithValues []*entities.Answer) map[int]int {
	result := make(map[int]int)
	for i := range answersWithValues {
		for j := range answersWithValues[i].Values {
			result[answersWithValues[i].Values[j].SkillId] += answersWithValues[i].Values[j].Points
		}
	}
	return result
}

func (a *addUserAnswers) getUserSkillsAndSkillChanges(skillsMap map[int]int, normalizationMap map[int]*entities.SkillNormalization, accountId int) ([]*entities.UserSkills, []*entities.SkillChange) {
	userSkills := make([]*entities.UserSkills, 0)
	skillChanges := make([]*entities.SkillChange, 0)
	for k, v := range skillsMap {
		normalization, exists := normalizationMap[k]
		if !exists {
			continue
		}

		points := int(float32(v-normalization.Min) / float32(normalization.Max-normalization.Min) * maxPoints)
		userSkills = append(userSkills, &entities.UserSkills{
			AccountId: accountId,
			SkillId:   k,
			Xp:        points,
		})

		skillChanges = append(skillChanges, &entities.SkillChange{
			AccountId: accountId,
			SkillId:   k,
			ActionId:  1, // todo сделать нормально (нужен ли action вообще?)
			Date:      time.Now().UTC(),
			Points:    points,
		})
	}

	return userSkills, skillChanges
}

func (a *addUserAnswers) getNewUserXpValue(userSkills []*entities.UserSkills) int {
	result := 0
	for i := range userSkills {
		result += userSkills[i].Xp
	}

	return result
}
