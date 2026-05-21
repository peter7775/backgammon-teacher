package dto

type HintResponse struct {
	Title          string   `json:"title"`
	Message        string   `json:"message"`
	Summary        string   `json:"summary"`
	Recommendation string   `json:"recommendation"`
	Classification string   `json:"classification"`
	BestMove       []string `json:"bestMove"`
}
