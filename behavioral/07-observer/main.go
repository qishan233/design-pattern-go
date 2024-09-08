package main

func main() {
	subject1 := NewSubject("subject1")
	subject2 := NewSubject("subject2")

	observer1 := NewObserver("observer1")
	observer2 := NewObserver("observer2")

	// build the relationship between observer and subject
	observer1.Observe(subject1)
	observer1.Observe(subject2)

	observer2.Observe(subject1)

	// subject change will notify correlative observers
	subject1.SetContent("content1")
	subject2.SetContent("content2")
}
