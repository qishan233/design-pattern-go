package main

import "fmt"

type Originator struct {
	CurrentState string
}

func (o *Originator) SetState(s string, needMemento bool) Memento {
	state := o.CurrentState
	o.CurrentState = s

	var m Memento
	if needMemento {
		m = NewMemento(state)
	}

	return m
}

func (o *Originator) CreateMemento() Memento {
	return NewMemento(o.CurrentState)
}

func (o *Originator) SetMemento(m Memento) {
	o.CurrentState = m.GetState()
}

func (o *Originator) Show() {
	fmt.Printf("Current State: %s\n", o.CurrentState)
}
