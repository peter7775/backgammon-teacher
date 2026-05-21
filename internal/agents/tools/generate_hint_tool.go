package tools

import (
	runtime "backgammon-teacher/internal/agents/runtime"
	analysis "backgammon-teacher/internal/modules/analysis/domain"
	coachapp "backgammon-teacher/internal/modules/coach/app"
)

type GenerateHintTool struct {
	Coach coachapp.GenerateHint
}

func (t GenerateHintTool) Name() string { return "generate_hint" }

func (t GenerateHintTool) Call(ctx runtime.TaskContext, input map[string]any) (map[string]any, error) {
	_ = ctx
	a, _ := input["analysis"].(analysis.PositionAnalysis)
	hint, err := t.Coach.Execute(a)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"title":          hint.Title,
		"message":        hint.Message,
		"summary":        hint.Summary,
		"recommendation": hint.Recommendation,
		"classification": hint.Classification,
		"bestMove":       hint.BestMove,
	}, nil
}
