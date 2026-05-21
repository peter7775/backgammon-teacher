package orchestration

import (
	runtime "backgammon-teacher/internal/agents/runtime"
	analysisapp "backgammon-teacher/internal/modules/analysis/app"
	coachapp "backgammon-teacher/internal/modules/coach/app"
	play "backgammon-teacher/internal/modules/play/domain"
)

type ReviewService struct {
	Analyze analysisapp.AnalyzeMove
	Coach   coachapp.GenerateHint
}

func (s ReviewService) ReviewMove(ctx runtime.TaskContext, pos play.Position, move play.Move) (map[string]any, error) {
	analysis, err := s.Analyze.Execute(pos, move)
	if err != nil {
		return nil, err
	}
	hint, err := s.Coach.Execute(analysis)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"summary":        hint.Summary,
		"recommendation": hint.Recommendation,
		"classification": hint.Classification,
		"bestMove":       hint.BestMove,
		"language":       ctx.Language,
	}, nil
}
