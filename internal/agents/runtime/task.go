package runtime

type Task struct {
	ID       string
	Goal     string
	Payload  map[string]any
	Metadata map[string]any
}

type TaskResult struct {
	Status   string
	Output   map[string]any
	Messages []string
}

type TaskContext struct {
	SessionID string
	UserID    string
	Language  string
}
