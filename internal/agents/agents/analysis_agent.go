package agents

import (
	runtime "backgammon-teacher/internal/agents/runtime"
	analysisapp "backgammon-teacher/internal/modules/analysis/app"
	play "backgammon-teacher/internal/modules/play/domain"
)

type AnalysisAgent struct {
	Analyze analysisapp.AnalyzeMove
}

func (a AnalysisAgent) Name() string { return "analysis-agent" }

func (a AnalysisAgent) Handle(ctx runtime.TaskContext, task runtime.Task) (runtime.TaskResult, error) {
	_ = ctx
	pos, _ := task.Payload["position"].(play.Position)
	move, _ := task.Payload["move"].(play.Move)
	result, err := a.Analyze.Execute(pos, move)
	if err != nil {
		return runtime.TaskResult{}, err
	}
	return runtime.TaskResult{
		Status: "ok",
		Output: map[string]any{
			"analysis": result,
		},
		Messages: []string{"move analyzed"},
	}, nil
}
