package reflection

type Critic struct{}

func (Critic) Review(output map[string]any) map[string]any {
	status := "ok"
	if output == nil || len(output) == 0 {
		status = "empty"
	}
	return map[string]any{"status": status}
}
