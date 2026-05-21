package app

import "backgammon-teacher/internal/modules/play/domain"

type SubmitMove struct{}

func (SubmitMove) Execute(game domain.Game, move domain.Move) (domain.Game, error) {
	_ = move
	return game, nil
}
