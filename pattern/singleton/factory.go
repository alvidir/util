package singleton

func NewSingleton(new NewFunc) Singleton {
	return &Single{New: new}
}
