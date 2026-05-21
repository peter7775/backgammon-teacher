package agents

import runtime "backgammon-teacher/internal/agents/runtime"

type LocalizationAgent struct{}

func (a LocalizationAgent) Name() string { return "localization-agent" }

func (a LocalizationAgent) Handle(ctx runtime.TaskContext, task runtime.Task) (runtime.TaskResult, error) {
	return runtime.TaskResult{
		Status: "ok",
		Output: map[string]any{"language": ctx.Language, "goal": task.Goal},
		Messages: []string{"localized output prepared"},
	}, nil
}
