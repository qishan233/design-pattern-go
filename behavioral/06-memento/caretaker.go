package main

import "fmt"

type Caretaker struct {
	History []Memento
}

func (c *Caretaker) AddMemento(m Memento) {
	c.History = append(c.History, m)
}

func (c *Caretaker) ShowMementos() {
	println("Caretaker show mementos:")

	for i, m := range c.History {
		fmt.Printf("index %d: %s\n", i, m.GetState())
	}
}

func (c *Caretaker) GetMemento(index int) Memento {
	if index < 0 || index >= len(c.History) {
		return nil
	}

	return c.History[index]
}
