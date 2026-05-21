package app

import (
	analysis "backgammon-teacher/internal/modules/analysis/domain"
	play "backgammon-teacher/internal/modules/play/domain"
)

type AnalyzeMove struct{}

func (AnalyzeMove) Execute(pos play.Position, move play.Move) (analysis.PositionAnalysis, error) {
	best := play.Move{Steps: []play.Step{{From: 13, To: 7, Pips: 6}, {From: 8, To: 7, Pips: 1}}}
	class := analysis.ClassificationBest
	summary := "Solid move."
	recommendation := "Keep building safely and reduce direct shots."

	if !sameMove(move, best) {
		class = analysis.ClassificationMistake
		summary = "Your move is playable, but it misses a stronger structural improvement."
		recommendation = "Prefer making a safer point when you can improve the board and reduce risk at the same time."
	}

	return analysis.PositionAnalysis{
		Position:   pos,
		PlayerMove: move,
		BestMove:   best,
		Candidates: []analysis.MoveScore{
			{Move: best, Equity: 0.15, EquityLoss: 0},
			{Move: move, Equity: 0.02, EquityLoss: 0.13},
		},
		Mistake: analysis.Mistake{Class: class, Themes: []analysis.Theme{analysis.ThemeSafety, analysis.ThemePrime}},
		Summary: summary,
		Recommendation: recommendation,
	}, nil
}

func sameMove(a, b play.Move) bool {
	if len(a.Steps) != len(b.Steps) {
		return false
	}
	for i := range a.Steps {
		if a.Steps[i] != b.Steps[i] {
			return false
		}
	}
	return true
}
