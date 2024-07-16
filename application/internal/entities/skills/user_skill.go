package skills

type UserSkill struct {
	SkillId int `json:"skill_id"`
	Xp      int `json:"points"`
}

type UserSkills struct {
	AccountId int          `json:"account_id"`
	Skills    []*UserSkill `json:"skills"`
}
