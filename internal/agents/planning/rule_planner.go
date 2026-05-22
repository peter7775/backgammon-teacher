package planning

import runtime "backgammon-teacher/internal/agents/runtime"

type RulePlanner struct{}

func (RulePlanner) BuildPlan(ctx runtime.TaskContext, task runtime.Task) (Plan, error) {
	_ = ctx
	switch task.Goal {
	case "review_move":
		return Plan{Steps: []PlanStep{
			{Agent: "analysis-agent", Goal: "analyze_move"},
			{Agent: "coach-agent", Goal: "generate_hint"},
			{Agent: "progress-agent", Goal: "update_progress"},
		}}, nil
	case "rag_answer":
		return Plan{Steps: []PlanStep{
			{Agent: "retrieval-agent", Goal: "retrieve_context"},
			{Agent: "coach-agent", Goal: "grounded_answer"},
		}}, nil
	case "recommend_lesson":
		return Plan{Steps: []PlanStep{{Agent: "lesson-agent", Goal: "recommend_lesson"}}}, nil
	default:
		return Plan{Steps: []PlanStep{{Agent: "coach-agent", Goal: task.Goal}}}, nil
	}
}
