package app

import (
	analysis "backgammon-teacher/internal/modules/analysis/domain"
	play "backgammon-teacher/internal/modules/play/domain"
)

type AnalyzeMove struct{}

func (AnalyzeMove) Execute(pos play.Position, move play.Move) (analysis.PositionAnalysis, error) {
	return analysis.PositionAnalysis{
		Position: pos,
		Candidates: []analysis.MoveScore{{Move: move, Equity: 0, EquityLoss: 0}},
		Mistake: analysis.Mistake{Class: analysis.ClassificationBest, Themes: []analysis.Theme{analysis.ThemeSafety}},
	}, nil
}
