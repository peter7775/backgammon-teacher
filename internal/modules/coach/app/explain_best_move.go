package app

import coach "backgammon-teacher/internal/modules/coach/domain"

type ExplainBestMove struct{}

func (ExplainBestMove) Execute() (coach.Explanation, error) {
	return coach.Explanation{Short: "Make the safer point.", Long: "This move improves structure and reduces the risk of leaving a direct shot."}, nil
}
