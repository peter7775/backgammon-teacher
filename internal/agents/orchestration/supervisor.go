package orchestration

import (
	"context"
	"backgammon-teacher/internal/agents/ports"
)

type Supervisor struct {
	runtime ports.Runtime
	memory ports.MemoryStore
}

func New(rt ports.Runtime, mem ports.MemoryStore) *Supervisor {
	return &Supervisor{runtime: rt, memory: mem}
}

func (s *Supervisor) Run(ctx context.Context, step ports.Step) error {
	return s.runtime.Execute(ctx, step)
}
