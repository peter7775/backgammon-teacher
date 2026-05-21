package app

import coach "backgammon-teacher/internal/modules/coach/domain"

type GenerateHint struct{}

func (GenerateHint) Execute(gameID string) (coach.Hint, error) {
	return coach.Hint{
		Title:   "Look for safety",
		Message: "Try to make a safer structure, reduce direct shots, and prefer moves that improve your board.",
	}, nil
}
