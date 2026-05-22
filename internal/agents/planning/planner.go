package planning

import "backgammon-teacher/internal/agents/ports"

func NextStep() ports.Step { return ports.Step("review_move") }
