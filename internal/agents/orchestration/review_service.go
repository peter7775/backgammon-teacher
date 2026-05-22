package orchestration

import coachapp "backgammon-teacher/internal/modules/coach/app"

type ReviewService struct {
	Generate coachapp.GenerateHintFunc
}

func NewReviewService(generate coachapp.GenerateHintFunc) *ReviewService {
	if generate == nil {
		generate = coachapp.DefaultGenerateHint
	}
	return &ReviewService{Generate: generate}
}
