package runtime

import (
	"fmt"

	memorypkg "backgammon-teacher/internal/agents/memory"
	planningpkg "backgammon-teacher/internal/agents/planning"
	toolspkg "backgammon-teacher/internal/agents/tools"
)

type Supervisor struct {
	Registry *Registry
	Planner  planningpkg.Planner
	Memory   memorypkg.Memory
	Tools    *toolspkg.Registry
}

func (s Supervisor) Execute(ctx TaskContext, task Task) (TaskResult, error) {
	plan, err := s.Planner.BuildPlan(ctx, task)
	if err != nil {
		return TaskResult{}, err
	}

	last := TaskResult{Status: "planned", Output: map[string]any{"steps": len(plan.Steps)}}
	for _, step := range plan.Steps {
		agent, ok := s.Registry.Get(step.Agent)
		if !ok {
			return TaskResult{}, fmt.Errorf("agent not found: %s", step.Agent)
		}
		res, err := agent.Handle(ctx, Task{
			ID:       task.ID,
			Goal:     step.Goal,
			Payload:  mergeMaps(task.Payload, step.Payload),
			Metadata: task.Metadata,
		})
		if err != nil {
			return TaskResult{}, err
		}
		_ = s.Memory.Append(ctx, memorypkg.Entry{Kind: "step_result", Content: res.Output})
		last = res
	}
	return last, nil
}

func mergeMaps(a, b map[string]any) map[string]any {
	out := map[string]any{}
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}
