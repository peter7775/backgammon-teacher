package domain

type HintGenerated struct { GameID string }
func (HintGenerated) Name() string { return "coach.hint_generated" }
