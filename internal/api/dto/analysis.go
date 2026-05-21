package dto

type AnalyzeMoveResponse struct {
	Summary        string   `json:"summary"`
	Recommendation string   `json:"recommendation"`
	Classification string   `json:"classification"`
	BestMove       []string `json:"bestMove"`
}
