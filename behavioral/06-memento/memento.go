package main

type Memento interface {
	GetState() string
}

type memento struct {
	state string
}

func (m *memento) GetState() string {
	return m.state
}

func NewMemento(value string) Memento {
	return &memento{state: value}
}
