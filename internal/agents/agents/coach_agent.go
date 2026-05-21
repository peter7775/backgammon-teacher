package agents

import (
	runtime "backgammon-teacher/internal/agents/runtime"
	analysis "backgammon-teacher/internal/modules/analysis/domain"
	coachapp "backgammon-teacher/internal/modules/coach/app"
)

type CoachAgent struct {
	Coach coachapp.GenerateHint
}

func (a CoachAgent) Name() string { return "coach-agent" }

func (a CoachAgent) Handle(ctx runtime.TaskContext, task runtime.Task) (runtime.TaskResult, error) {
	_ = ctx
	an, _ := task.Payload["analysis"].(analysis.PositionAnalysis)
	hint, err := a.Coach.Execute(an)
	if err != nil {
		return runtime.TaskResult{}, err
	}
	return runtime.TaskResult{
		Status: "ok",
		Output: map[string]any{"hint": hint},
		Messages: []string{"hint generated"},
	}, nil
}
