package ports

import learning "backgammon-teacher/internal/modules/learning/domain"

type LessonRepository interface {
	Get(id string) (learning.Lesson, error)
}
