package domain

type MoveAnalyzed struct { GameID string }
func (MoveAnalyzed) Name() string { return "analysis.move_analyzed" }
