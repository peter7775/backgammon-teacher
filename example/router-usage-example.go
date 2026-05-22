package main

import (
	"fmt"
	"backgammon-teacher/llm"
)

func main() {
	primary := map[llm.UseCase]llm.Provider{
		llm.UseCasePlanning:      llm.NewClaudeProvider("claude-opus-4-7"),
		llm.UseCaseHardReasoning:  llm.NewClaudeProvider("claude-opus-4-7"),
		llm.UseCaseReviewMove:     llm.NewClaudeProvider("claude-sonnet-4-6"),
		llm.UseCaseRAGAnswer:      llm.NewClaudeProvider("claude-sonnet-4-6"),
	}
	fallbacks := map[llm.UseCase]llm.Provider{
		llm.UseCaseOffline: llm.NewOllamaProvider("gpt-oss-20b"),
	}

	router := llm.NewRouter(primary, fallbacks)
	resp, err := router.Route(llm.Request{UseCase: llm.UseCasePlanning, Prompt: "Plan the next teaching step"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Model)
	fmt.Println(resp.Content)
}
