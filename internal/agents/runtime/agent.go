package runtime

type Agent interface {
	Name() string
	Handle(ctx TaskContext, task Task) (TaskResult, error)
}
