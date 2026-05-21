package planning

import runtime "backgammon-teacher/internal/agents/runtime"

type Planner interface {
	BuildPlan(ctx runtime.TaskContext, task runtime.Task) (Plan, error)
}

type Plan struct {
	Steps []PlanStep
}

type PlanStep struct {
	Agent   string
	Goal    string
	Payload map[string]any
}
