package app

import (
	"backgammon-teacher/internal/modules/play/domain"
	"backgammon-teacher/internal/modules/play/ports"
)

type GetGame struct {
	Games ports.GameRepository
}

func (uc GetGame) Execute(gameID string) (domain.Game, error) {
	return uc.Games.Get(gameID)
}
