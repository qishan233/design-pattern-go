package main

import "fmt"

type Observer interface {
	Update(Subject)
	Observe(Subject)
}

type ConcreteObserver struct {
	Name string
}

func NewObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{Name: name}
}

func (co *ConcreteObserver) Observe(s Subject) {
	s.Attach(co)
}
func (co *ConcreteObserver) Update(s Subject) {
	fmt.Printf("%s received the message from %s, content: %s\n", co.Name, s.GetID(), s.GetContent())
}
