package adapter

// New builds a new Encoder
func New(marshal Marshal, unmarshal Unmarshal) Adapter {
	return &Adaptee{X_Marshal: marshal, X_Unmarshal: unmarshal}
}
