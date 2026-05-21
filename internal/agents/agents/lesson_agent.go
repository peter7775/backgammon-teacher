package agents

import (
	runtime "backgammon-teacher/internal/agents/runtime"
	learningapp "backgammon-teacher/internal/modules/learning/app"
)

type LessonAgent struct {
	Recommend learningapp.RecommendNextLesson
}

func (a LessonAgent) Name() string { return "lesson-agent" }

func (a LessonAgent) Handle(ctx runtime.TaskContext, task runtime.Task) (runtime.TaskResult, error) {
	_ = ctx
	_ = task
	rec, err := a.Recommend.Execute()
	if err != nil {
		return runtime.TaskResult{}, err
	}
	return runtime.TaskResult{Status: "ok", Output: map[string]any{"lesson": rec}, Messages: []string{"lesson recommended"}}, nil
}
