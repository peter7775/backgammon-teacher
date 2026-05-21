package app

import (
	"backgammon-teacher/internal/modules/play/domain"
	"backgammon-teacher/internal/modules/play/ports"
)

type StartGame struct {
	Games ports.GameRepository
}

func (uc StartGame) Execute(gameID string) (domain.Game, error) {
	game := domain.StartGame(gameID)
	if err := uc.Games.Save(game); err != nil {
		return domain.Game{}, err
	}
	return game, nil
}
