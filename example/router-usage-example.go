package main

import (
	"backgammon-teacher/internal/llm"
	"fmt"
)

func main() {
	router := llm.NewRouter(
		map[llm.UseCase]llm.Provider{
			llm.UseCasePlanning: llm.NewClaudeProvider("claude-opus-4-7"),
		},
		map[llm.UseCase]llm.Provider{
			llm.UseCaseOffline: llm.NewOllamaProvider("gpt-oss-20b"),
		},
	)
	resp, _ := router.Route(llm.Request{UseCase: llm.UseCasePlanning, Prompt: "Plan the next teaching step"})
	fmt.Println(resp.Model)
}
