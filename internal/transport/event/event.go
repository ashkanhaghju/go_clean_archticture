package event

type Event interface {
	StartConsume()
	StopConsume()
}
