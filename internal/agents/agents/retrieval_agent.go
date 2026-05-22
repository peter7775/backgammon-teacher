package agents

import (
	runtime "backgammon-teacher/internal/agents/runtime"
	knowledge "backgammon-teacher/internal/modules/knowledge/domain"
	knowledgeapp "backgammon-teacher/internal/modules/knowledge/app"
)

type RetrievalAgent struct {
	Retrieve knowledgeapp.RetrieveContext
}

func (a RetrievalAgent) Name() string { return "retrieval-agent" }

func (a RetrievalAgent) Handle(ctx runtime.TaskContext, task runtime.Task) (runtime.TaskResult, error) {
	query, _ := task.Payload["query"].(string)
	matches, err := a.Retrieve.Execute(knowledge.Query{
		Text:      query,
		Language:  ctx.Language,
		UserID:    ctx.UserID,
		SessionID: ctx.SessionID,
		TopK:      5,
	})
	if err != nil {
		return runtime.TaskResult{}, err
	}
	return runtime.TaskResult{
		Status: "ok",
		Output: map[string]any{"matches": matches},
		Messages: []string{"context retrieved from local and shared knowledge bases"},
	}, nil
}
