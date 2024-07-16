package usecases

import (
	"context"
	"smartri_app/controllers/requests"
	"smartri_app/internal"
	"smartri_app/internal/entities/skills_entities"
	"smartri_app/internal/entities/test_entities"
	"smartri_app/internal/entities/user_data_entities"
	"smartri_app/internal/errs"
	"time"
)

const maxPoints = 200

type addUserAnswers struct {
	testRepo        internal.ITestRepository
	skillRepo       internal.ISkillRepository
	userDataRepo    internal.IUserDataRepository
	userAnswersRepo internal.IUserAnswersRepository
}

func NewAddUserAnswers(
	testRepo internal.ITestRepository,
	skillRepo internal.ISkillRepository,
	userDataRepo internal.IUserDataRepository,
	userAnswersRepo internal.IUserAnswersRepository) internal.IAddUserTestAnswersUseCase {
	return &addUserAnswers{testRepo: testRepo, skillRepo: skillRepo, userDataRepo: userDataRepo, userAnswersRepo: userAnswersRepo}
}

func (a *addUserAnswers) Add(context context.Context, answers *requests.UserAnswersRequest, accountId int) (*requests.UserAnswersResponse, error) {
	userData, err := a.userDataRepo.GetByAccountId(context, accountId)

	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, errs.UserDataNotFoundError
	}
	userHasAnswers, err := a.userAnswersRepo.CheckUserHasAnswers(context, accountId)

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
	normalizationBySkillIdMap, err := a.getNormalizationMap(context, skills)
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
	userData.XP = user_data_entities.XP(a.getNewUserXpValue(userSkills))

	err = a.testRepo.AddUserAnswersWithSkillChanges(context, entityUserAnswers, skillChanges, userSkills, userData)

	if err != nil {
		return nil, err
	}

	result, err := a.skillRepo.GetSkillsByAccountId(context, accountId)

	if err != nil {
		return nil, err
	}

	return &requests.UserAnswersResponse{
		AccountId: accountId,
		Skills:    result.Skills,
		TotalExp:  int(userData.XP),
	}, nil
}

func (a *addUserAnswers) getEntityAnswers(accountId int, answers *requests.UserAnswersRequest) *test_entities.UserTestAnswers {
	result := &test_entities.UserTestAnswers{AccountId: accountId, Answers: make([]test_entities.UserTestAnswer, 0)}
	for i := range answers.Answers {
		result.Answers = append(result.Answers, test_entities.UserTestAnswer{
			QuestionId: answers.Answers[i].QuestionId,
			AnswerId:   answers.Answers[i].AnswerId,
		})
	}

	return result
}

func (a *addUserAnswers) getNormalizationMap(context context.Context, skills []*skills_entities.Skill) (map[int]*skills_entities.SkillNormalization, error) {
	result := make(map[int]*skills_entities.SkillNormalization)
	for i := range skills {
		r, err := a.skillRepo.GetSkillNormalizationBySkillId(context, skills[i].Id)
		result[i] = r
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (a *addUserAnswers) getAnswersWithValues(context context.Context, entityAnswers *test_entities.UserTestAnswers) ([]*test_entities.Answer, error) {
	result := make([]*test_entities.Answer, 0)
	for i := 0; i < len(entityAnswers.Answers); i++ {
		answer, err := a.testRepo.GetAnswerWithValues(context, entityAnswers.Answers[i].AnswerId)
		if err != nil {
			return nil, err
		}

		result = append(result, answer)
	}

	return result, nil
}

func (a *addUserAnswers) getSkillIdPointsMap(answersWithValues []*test_entities.Answer) map[int]int {
	result := make(map[int]int)
	for i := range answersWithValues {
		for j := range answersWithValues[i].Values {
			result[answersWithValues[i].Values[j].SkillId] += answersWithValues[i].Values[j].Points
		}
	}
	return result
}

func (a *addUserAnswers) getUserSkillsAndSkillChanges(skillsMap map[int]int, normalizationMap map[int]*skills_entities.SkillNormalization, accountId int) (*skills_entities.UserSkills, []*skills_entities.SkillChange) {
	skills := make([]*skills_entities.UserSkill, 0)
	skillChanges := make([]*skills_entities.SkillChange, 0)
	for k, v := range skillsMap {
		normalization, exists := normalizationMap[k]
		if !exists {
			continue
		}

		points := int(float32(v-normalization.Min) / float32(normalization.Max-normalization.Min) * maxPoints)
		skills = append(skills, &skills_entities.UserSkill{
			SkillId: k,
			Xp:      points,
		})

		skillChanges = append(skillChanges, &skills_entities.SkillChange{
			AccountId: accountId,
			SkillId:   k,
			ActionId:  1, // todo сделать нормально (нужен ли action вообще?)
			Date:      time.Now().UTC(),
			Points:    points,
		})
	}

	userSkills := skills_entities.UserSkills{
		AccountId: accountId,
		Skills:    skills,
	}
	return &userSkills, skillChanges
}

func (a *addUserAnswers) getNewUserXpValue(userSkills *skills_entities.UserSkills) int {
	result := 0
	for i := range userSkills.Skills {
		result += userSkills.Skills[i].Xp
	}

	return result
}
