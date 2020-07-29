package observer

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	Broadcast(interface{})
}

type Observer interface {
	OnUpdate(interface{})
}
