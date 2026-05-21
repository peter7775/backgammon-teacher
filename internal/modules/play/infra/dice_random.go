package infra

import (
	"math/rand"
	"time"

	"backgammon-teacher/internal/modules/play/domain"
)

type RandomDice struct{}

func (RandomDice) Roll() domain.Dice {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return domain.Dice{A: int8(r.Intn(6) + 1), B: int8(r.Intn(6) + 1)}
}
