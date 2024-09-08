package main

func main() {
	caretaker := &Caretaker{
		History: make([]Memento, 0),
	}

	originator := &Originator{CurrentState: "initial state"}

	// get memento that represent the last state of originator and store to caretaker
	m1 := originator.SetState("state 1", true)
	caretaker.AddMemento(m1)

	m2 := originator.SetState("state 2", true)
	caretaker.AddMemento(m2)

	m3 := originator.SetState("state 3", true)
	caretaker.AddMemento(m3)

	// print the current state of originator
	originator.Show()

	// print all mementos in caretaker
	caretaker.ShowMementos()

	// use the memento stored in caretaker to restore the state of originator
	t := caretaker.GetMemento(2)
	originator.SetMemento(t)

	originator.Show()
}
