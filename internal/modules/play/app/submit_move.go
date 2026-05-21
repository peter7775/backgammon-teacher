package app

import (
	"backgammon-teacher/internal/modules/play/domain"
	"backgammon-teacher/internal/modules/play/ports"
)

type SubmitMove struct {
	Games ports.GameRepository
}

func (uc SubmitMove) Execute(gameID string, move domain.Move) (domain.Game, error) {
	game, err := uc.Games.Get(gameID)
	if err != nil {
		return domain.Game{}, err
	}
	game = game.ApplyMove(move)
	if err := uc.Games.Save(game); err != nil {
		return domain.Game{}, err
	}
	return game, nil
}
