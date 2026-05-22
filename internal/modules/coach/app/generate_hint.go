package app

import "strings"

type GenerateHintFunc func(HintInput) HintOutput

func DefaultGenerateHint(in HintInput) HintOutput {
	title := "Good move"
	summary := "The played move is close to optimal."
	suggestion := "Continue focusing on position safety and timing."
	explanation := ""

	switch strings.ToLower(in.Classification) {
	case "best":
		title = "Best move"
		summary = "You played the optimal move."
		suggestion = "Keep reinforcing the same decision pattern."
	case "inaccuracy":
		title = "Inaccuracy"
		summary = "The move is reasonable, but not quite best."
		suggestion = "Compare the chosen move with the optimal alternative."
	case "mistake":
		title = "Mistake"
		summary = "The move loses noticeable equity."
		suggestion = "Review why the optimal move is stronger in this position."
	case "blunder":
		title = "Blunder"
		summary = "The move loses a lot of equity."
		suggestion = "Stop and analyze the tactical and positional consequences."
	default:
		if in.Classification != "" {
			title = strings.ToUpper(in.Classification[:1]) + strings.ToLower(in.Classification[1:])
			summary = "The move has been analyzed."
			suggestion = "Use the optimal alternative as a reference."
		}
	}

	if len(in.Notes) > 0 {
		explanation = strings.Join(in.Notes, " ")
	}
	if explanation == "" {
		explanation = summary
	}

	return HintOutput{Title: title, Summary: summary, Suggestion: suggestion, Explanation: explanation}
}
