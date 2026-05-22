package app

import "backgammon-teacher/internal/modules/play/domain"

func MapAnalysisToHintInput(a domain.Analysis) HintInput {
	return HintInput{
		PositionID:     a.PositionID,
		PlayedMove:     a.PlayedMove,
		OptimalMove:    a.OptimalMove,
		EquityLoss:     a.EquityLoss,
		Classification: a.Classification,
		Notes:          append([]string(nil), a.Notes...),
	}
}
