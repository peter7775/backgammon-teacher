package ports

import "backgammon-teacher/internal/modules/play/domain"

type GameRepository interface {
	Save(game domain.Game) error
	Get(id string) (domain.Game, error)
}
