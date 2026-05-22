package memory

import (
	"context"
	"sync"

	"backgammon-teacher/internal/agents/ports"
)

type InMemoryStore struct {
	mu sync.RWMutex
	data map[string]any
}

func New() *InMemoryStore { return &InMemoryStore{data: map[string]any{}} }

func (s *InMemoryStore) Get(ctx context.Context, key string) (any, bool, error) {
	s.mu.RLock(); defer s.mu.RUnlock()
	v, ok := s.data[key]
	return v, ok, nil
}

func (s *InMemoryStore) Set(ctx context.Context, key string, value any) error {
	s.mu.Lock(); defer s.mu.Unlock()
	s.data[key] = value
	return nil
}

var _ ports.MemoryStore = (*InMemoryStore)(nil)
