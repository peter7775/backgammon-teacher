package memory

import types "backgammon-teacher/internal/agents/types"

type Memory interface {
	Append(ctx types.TaskContext, entry Entry) error
	Recent(ctx types.TaskContext, limit int) ([]Entry, error)
}

type Entry struct {
	Kind    string
	Content map[string]any
}
