package domain

import play "backgammon-teacher/internal/modules/play/domain"

type PositionAnalysis struct {
	Position        play.Position
	PlayerMove      play.Move
	BestMove        play.Move
	Candidates      []MoveScore
	Mistake         Mistake
	Summary         string
	Recommendation  string
}
