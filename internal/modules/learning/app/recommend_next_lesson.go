package app

import learning "backgammon-teacher/internal/modules/learning/domain"

type RecommendNextLesson struct{}

func (RecommendNextLesson) Execute() (learning.Recommendation, error) {
	return learning.Recommendation{LessonID: "safe-play-101", Reason: "Repeated safety mistakes."}, nil
}
