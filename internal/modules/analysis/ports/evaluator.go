package ports

import (
	analysis "backgammon-teacher/internal/modules/analysis/domain"
	play "backgammon-teacher/internal/modules/play/domain"
)

type Evaluator interface {
	Analyze(pos play.Position, move play.Move) (analysis.PositionAnalysis, error)
}
