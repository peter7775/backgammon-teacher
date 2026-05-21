package memory

import runtime "backgammon-teacher/internal/agents/runtime"

type InMemoryStore struct {
	items map[string][]Entry
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{items: map[string][]Entry{}}
}

func (s *InMemoryStore) Append(ctx runtime.TaskContext, entry Entry) error {
	key := ctx.UserID + ":" + ctx.SessionID
	s.items[key] = append(s.items[key], entry)
	return nil
}

func (s *InMemoryStore) Recent(ctx runtime.TaskContext, limit int) ([]Entry, error) {
	key := ctx.UserID + ":" + ctx.SessionID
	entries := s.items[key]
	if len(entries) <= limit {
		return entries, nil
	}
	return entries[len(entries)-limit:], nil
}
