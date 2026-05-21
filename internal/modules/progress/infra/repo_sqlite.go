package infra

import (
	progress "backgammon-teacher/internal/modules/progress/domain"
	"backgammon-teacher/internal/platform/errors"
)

type SQLiteProfileRepository struct {
	items map[string]progress.LearnerProfile
}

func NewSQLiteProfileRepository() *SQLiteProfileRepository {
	return &SQLiteProfileRepository{items: map[string]progress.LearnerProfile{}}
}

func (r *SQLiteProfileRepository) Save(profile progress.LearnerProfile) error {
	r.items[profile.UserID] = profile
	return nil
}

func (r *SQLiteProfileRepository) Get(userID string) (progress.LearnerProfile, error) {
	p, ok := r.items[userID]
	if !ok {
		return progress.LearnerProfile{}, errors.ErrNotFound
	}
	return p, nil
}
