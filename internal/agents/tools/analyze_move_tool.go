package tools

import (
	runtime "backgammon-teacher/internal/agents/runtime"
	analysisapp "backgammon-teacher/internal/modules/analysis/app"
	play "backgammon-teacher/internal/modules/play/domain"
)

type AnalyzeMoveTool struct {
	Analyze analysisapp.AnalyzeMove
}

func (t AnalyzeMoveTool) Name() string { return "analyze_move" }

func (t AnalyzeMoveTool) Call(ctx runtime.TaskContext, input map[string]any) (map[string]any, error) {
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
