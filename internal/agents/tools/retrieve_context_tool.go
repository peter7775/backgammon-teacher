package tools

import (
	types "backgammon-teacher/internal/agents/types"
	knowledge "backgammon-teacher/internal/modules/knowledge/domain"
	knowledgeapp "backgammon-teacher/internal/modules/knowledge/app"
)

type RetrieveContextTool struct {
	Retrieve knowledgeapp.RetrieveContext
}

func (t RetrieveContextTool) Name() string { return "retrieve_context" }

func (t RetrieveContextTool) Call(ctx types.TaskContext, input map[string]any) (map[string]any, error) {
	query, _ := input["query"].(string)
	matches, err := t.Retrieve.Execute(knowledge.Query{
		Text:      query,
		Language:  ctx.Language,
		UserID:    ctx.UserID,
		SessionID: ctx.SessionID,
		TopK:      5,
	})
	if err != nil {
		return nil, err
	}
	return map[string]any{"matches": matches}, nil
}
