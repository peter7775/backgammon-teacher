package agents

import coachapp "backgammon-teacher/internal/modules/coach/app"

type CoachAgent struct {
	Generate coachapp.GenerateHintFunc
}

func NewCoachAgent(generate coachapp.GenerateHintFunc) *CoachAgent {
	if generate == nil {
		generate = coachapp.DefaultGenerateHint
	}
	return &CoachAgent{Generate: generate}
}
