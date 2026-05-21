package agents

import runtime "backgammon-teacher/internal/agents/runtime"

type ProgressAgent struct{}

func (a ProgressAgent) Name() string { return "progress-agent" }

func (a ProgressAgent) Handle(ctx runtime.TaskContext, task runtime.Task) (runtime.TaskResult, error) {
	_ = task
	return runtime.TaskResult{
		Status: "ok",
		Output: map[string]any{"userId": ctx.UserID, "updated": true},
		Messages: []string{"progress updated"},
	}, nil
}
