package memory

import runtime "backgammon-teacher/internal/agents/runtime"

type Memory interface {
	Append(ctx runtime.TaskContext, entry Entry) error
	Recent(ctx runtime.TaskContext, limit int) ([]Entry, error)
}

type Entry struct {
	Kind    string
	Content map[string]any
}
