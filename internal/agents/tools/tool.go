package tools

import runtime "backgammon-teacher/internal/agents/runtime"

type Tool interface {
	Name() string
	Call(ctx runtime.TaskContext, input map[string]any) (map[string]any, error)
}
