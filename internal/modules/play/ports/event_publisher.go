package ports

type EventPublisher interface {
	Publish(events ...any) error
}
