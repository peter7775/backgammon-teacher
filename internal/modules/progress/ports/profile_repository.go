package ports

import progress "backgammon-teacher/internal/modules/progress/domain"

type ProfileRepository interface {
	Save(profile progress.LearnerProfile) error
	Get(userID string) (progress.LearnerProfile, error)
}
