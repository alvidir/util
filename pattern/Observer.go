package pattern

// An Observer represents an object that spects some Subject notifications
type Observer interface {
	OnNotification(interface{})
}

// A Subject represents an object that's being observed by an Observer
type Subject interface {
	AddObserver(Observer)
	RemoveObserver(Observer)

	Broadcast(interface{})
}
