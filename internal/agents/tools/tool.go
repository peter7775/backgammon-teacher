package tools

import types "backgammon-teacher/internal/agents/types"

type Tool interface {
	Name() string
	Call(ctx types.TaskContext, input map[string]any) (map[string]any, error)
}
