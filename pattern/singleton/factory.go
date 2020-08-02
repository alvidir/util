package singleton

func NewSingleton(new NewFunc) Singleton {
	return &singleton{New: new}
}
