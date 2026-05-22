package runtime

import (
	"context"
	"backgammon-teacher/internal/agents/ports"
)

type BasicRuntime struct{}

func New() *BasicRuntime { return &BasicRuntime{} }

func (r *BasicRuntime) Execute(ctx context.Context, step ports.Step) error { return nil }

var _ ports.Runtime = (*BasicRuntime)(nil)
