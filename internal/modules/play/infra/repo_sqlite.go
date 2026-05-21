package infra

import (
	"backgammon-teacher/internal/modules/play/domain"
	"backgammon-teacher/internal/platform/errors"
)

type SQLiteGameRepository struct {
	items map[string]domain.Game
}

func NewSQLiteGameRepository() *SQLiteGameRepository {
	return &SQLiteGameRepository{items: map[string]domain.Game{}}
}

func (r *SQLiteGameRepository) Save(game domain.Game) error {
	r.items[game.ID] = game
	return nil
}

func (r *SQLiteGameRepository) Get(id string) (domain.Game, error) {
	g, ok := r.items[id]
	if !ok {
		return domain.Game{}, errors.ErrNotFound
	}
	return g, nil
}
