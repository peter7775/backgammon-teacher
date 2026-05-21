package ports

import coach "backgammon-teacher/internal/modules/coach/domain"

type ExplanationRenderer interface {
	RenderHint() (coach.Hint, error)
	RenderExplanation() (coach.Explanation, error)
}
