package domain

import "backgammon-teacher/internal/modules/play/domain"

type MoveScore struct {
	Move       domain.Move
	Equity     float64
	EquityLoss float64
}
