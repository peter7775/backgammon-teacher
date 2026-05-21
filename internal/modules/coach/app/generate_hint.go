package app

import coach "backgammon-teacher/internal/modules/coach/domain"

type GenerateHint struct{}

func (GenerateHint) Execute() (coach.Hint, error) {
	return coach.Hint{Title: "Look for safety", Message: "Try to improve your board while reducing direct shots."}, nil
}
