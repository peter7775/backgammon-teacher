package planning

import types "backgammon-teacher/internal/agents/types"

// Planner is the interface used by the runtime to build execution plans.
type Planner interface {
	BuildPlan(ctx types.TaskContext, task types.Task) (types.Plan, error)
}

type RulePlanner struct{}

func (RulePlanner) BuildPlan(ctx types.TaskContext, task types.Task) (types.Plan, error) {
	_ = ctx
	switch task.Goal {
	case "review_move":
		return types.Plan{Steps: []types.PlanStep{
			{Agent: "analysis-agent", Goal: "analyze_move"},
			{Agent: "coach-agent", Goal: "generate_hint"},
			{Agent: "progress-agent", Goal: "update_progress"},
		}}, nil
	case "rag_answer":
		return types.Plan{Steps: []types.PlanStep{
			{Agent: "retrieval-agent", Goal: "retrieve_context"},
			{Agent: "coach-agent", Goal: "grounded_answer"},
		}}, nil
	case "recommend_lesson":
		return types.Plan{Steps: []types.PlanStep{{Agent: "lesson-agent", Goal: "recommend_lesson"}}}, nil
	default:
		return types.Plan{Steps: []types.PlanStep{{Agent: "coach-agent", Goal: task.Goal}}}, nil
	}
}
