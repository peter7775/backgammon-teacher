package infra

import coach "backgammon-teacher/internal/modules/coach/domain"

type TemplateRenderer struct{}

func (TemplateRenderer) RenderHint() (coach.Hint, error) {
	return coach.Hint{Title: "Template hint", Message: "Build a stronger board before racing."}, nil
}

func (TemplateRenderer) RenderExplanation() (coach.Explanation, error) {
	return coach.Explanation{Short: "Safer move.", Long: "This line follows a safer teaching principle."}, nil
}
