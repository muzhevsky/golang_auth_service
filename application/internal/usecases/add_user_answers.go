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
	hasAnswers, err := a.userRepo.CheckUserHasAnswers(context, accountId)
	if err != nil {
		return nil, err
	}

	if hasAnswers {
		return nil, errs.UserHasAlreadyPassedTestError
	}

	entityAnswers := &entities.UserTestAnswers{accountId, make([]entities.UserTestAnswer, 0)}

	userData, err := a.userRepo.GetDataByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, errs.UserDataNotFoundError
	}

	skills, err := a.skillRepo.GetAllSkills(context)
	if err != nil {
		return nil, err
	}

	normalizationMap := make(map[int]*entities.SkillNormalization)
	for i := range skills {
		normalizationMap[i], err = a.skillRepo.GetSkillNormalizationBySkillId(context, skills[i].Id)
		if err != nil {
			return nil, err
		}
	}

	for i := range answers.Answers {
		entityAnswers.Answers = append(entityAnswers.Answers, entities.UserTestAnswer{
			QuestionId: answers.Answers[i].QuestionId,
			AnswerId:   answers.Answers[i].AnswerId,
		})
	}

	answersWithValues := make([]*entities.Answer, 0)
	for i := 0; i < len(entityAnswers.Answers); i++ {
		answer, err := a.testRepo.GetAnswerWithValues(context, entityAnswers.Answers[i].AnswerId)
		if err != nil {
			return nil, err
		}

		answersWithValues = append(answersWithValues, answer)
	}

	skillsMap := make(map[int]int)
	for i := range answersWithValues {
		for j := range answersWithValues[i].Values {
			skillsMap[answersWithValues[i].Values[j].SkillId] += answersWithValues[i].Values[j].Points
		}
	}

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
			ActionId:  1, // todo сделать нормально
			Date:      time.Now().UTC(),
			Points:    points,
		})
	}

	newXP := 0
	for i := range userSkills {
		newXP += userSkills[i].Xp
	}

	userData.XP = newXP

	err = a.testRepo.AddUserTestResults(context, entityAnswers, skillChanges, userSkills, userData)
	if err != nil {
		return nil, err
	}

	result, err := a.skillRepo.GetSkillsByAccountId(context, accountId)
	if err != nil {
		return nil, err
	}

	for i := range result {
		result[i].AccountId = 0
	}

	return &requests.UserAnswersResponse{
		AccountId: accountId,
		Skills:    result,
		TotalExp:  newXP,
	}, nil
}
