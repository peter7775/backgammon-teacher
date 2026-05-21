package domain

import play "backgammon-teacher/internal/modules/play/domain"

type PositionAnalysis struct {
	Position   play.Position
	Candidates []MoveScore
	Mistake    Mistake
}
