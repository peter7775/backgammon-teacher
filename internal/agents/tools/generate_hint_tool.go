package tools

import coachapp "backgammon-teacher/internal/modules/coach/app"

type GenerateHintTool struct {
	Generate coachapp.GenerateHintFunc
}

func NewGenerateHintTool(generate coachapp.GenerateHintFunc) *GenerateHintTool {
	if generate == nil {
		generate = coachapp.DefaultGenerateHint
	}
	return &GenerateHintTool{Generate: generate}
}
