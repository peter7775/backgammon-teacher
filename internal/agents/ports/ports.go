package ports

import "context"

type Step string

type MemoryStore interface {
	Get(ctx context.Context, key string) (any, bool, error)
	Set(ctx context.Context, key string, value any) error
}

type Runtime interface {
	Execute(ctx context.Context, step Step) error
}
