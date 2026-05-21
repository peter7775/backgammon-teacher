package ports

import "backgammon-teacher/internal/modules/play/domain"

type RandomDice interface {
	Roll() domain.Dice
}
