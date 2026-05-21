package bus

type Event interface {
	Name() string
}

type Publisher interface {
	Publish(events ...Event) error
}
