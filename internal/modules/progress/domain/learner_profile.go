package domain

type LearnerProfile struct {
	UserID   string
	Skills   []SkillProgress
	Weakness []Weakness
}
