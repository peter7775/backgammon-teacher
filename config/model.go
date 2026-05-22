package config

type ModelConfig struct {
	Primary ModelProvider `yaml:"primary" json:"primary"`
	Default ModelProvider `yaml:"default" json:"default"`
	Local   ModelProvider `yaml:"local" json:"local"`
	Routing ModelRouting  `yaml:"routing" json:"routing"`
}

type ModelProvider struct {
	Name       string   `yaml:"name" json:"name"`
	Model      string   `yaml:"model" json:"model"`
	Role       string   `yaml:"role" json:"role"`
	UseCases   []string `yaml:"use_cases" json:"use_cases"`
	FallbackTo string   `yaml:"fallback_to" json:"fallback_to"`
}

type ModelRouting struct {
	ReviewMove    string `yaml:"review_move" json:"review_move"`
	HardReasoning string `yaml:"hard_reasoning" json:"hard_reasoning"`
	RAGAnswer     string `yaml:"rag_answer" json:"rag_answer"`
	Planning      string `yaml:"planning" json:"planning"`
	Offline       string `yaml:"offline" json:"offline"`
}
