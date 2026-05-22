package tools

import "backgammon-teacher/internal/agents/ports"

func Validate(step ports.Step) bool { return step != "" }
