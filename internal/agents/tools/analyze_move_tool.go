package tools

import (
	types "backgammon-teacher/internal/agents/types"
	analysisapp "backgammon-teacher/internal/modules/analysis/app"
	play "backgammon-teacher/internal/modules/play/domain"
)

type AnalyzeMoveTool struct {
	Analyze analysisapp.AnalyzeMove
}

func (t AnalyzeMoveTool) Name() string { return "analyze_move" }

func (t AnalyzeMoveTool) Call(ctx types.TaskContext, input map[string]any) (map[string]any, error) {
	_ = ctx
	pos, _ := input["position"].(play.Position)
	move, _ := input["move"].(play.Move)
	result, err := t.Analyze.Execute(pos, move)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"summary":        result.Summary,
		"recommendation": result.Recommendation,
		"classification": string(result.Mistake.Class),
	}, nil
}
