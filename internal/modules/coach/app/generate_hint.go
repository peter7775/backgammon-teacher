package app

import (
	"fmt"

	analysis "backgammon-teacher/internal/modules/analysis/domain"
	coach "backgammon-teacher/internal/modules/coach/domain"
)

type GenerateHint struct{}

func (GenerateHint) Execute(a analysis.PositionAnalysis) (coach.Hint, error) {
	best := make([]string, 0, len(a.BestMove.Steps))
	for _, st := range a.BestMove.Steps {
		best = append(best, fmt.Sprintf("%d/%d", st.From, st.To))
	}

	title := "Good move"
	message := "You chose a strong move. Keep focusing on safety and structure."
	if a.Mistake.Class != analysis.ClassificationBest {
		title = "Try a safer plan"
		message = "There was a stronger move available. Look for a move that improves your board and lowers tactical risk."
	}

	return coach.Hint{
		Title:          title,
		Message:        message,
		Summary:        a.Summary,
		Recommendation: a.Recommendation,
		BestMove:       best,
		Classification: string(a.Mistake.Class),
	}, nil
}
