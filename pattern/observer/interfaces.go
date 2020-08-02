package observer

// Subject represents an element that will notify to a set of observers the happening of something
type Subject interface {
	Register(Observer)
	Unregister(Observer)
	Broadcast(interface{})
}

// Observer represents an object waiting for notifications from a subject
type Observer interface {
	OnNotification(interface{})
}
