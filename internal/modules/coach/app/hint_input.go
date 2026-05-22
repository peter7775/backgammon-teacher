package app

type HintInput struct {
	PositionID     string
	PlayedMove     string
	OptimalMove    string
	EquityLoss     float64
	Classification string
	Notes          []string
}

type HintOutput struct {
	Title       string
	Summary     string
	Suggestion  string
	Explanation string
}
