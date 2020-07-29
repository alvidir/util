package observer

// NewSubject builds a brand new and empty subject
func NewSubject() Subject {
	return &subject{}
}
