package main

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()

	// interface for observer to get the content of subject
	GetContent() string
	GetID() string

	SetContent(string)
}

type ConcreteSubject struct {
	Observers []Observer
	Content   string
	ID        string
}

func NewSubject(id string) Subject {
	return &ConcreteSubject{ID: id}
}

func (cs *ConcreteSubject) Attach(o Observer) {
	cs.Observers = append(cs.Observers, o)
}

func (cs *ConcreteSubject) Detach(o Observer) {
	for i, observer := range cs.Observers {
		if observer == o {
			cs.Observers = append(cs.Observers[:i], cs.Observers[i+1:]...)
			break
		}
	}
}

func (cs *ConcreteSubject) Notify() {
	for _, observer := range cs.Observers {
		observer.Update(cs)
	}
}

func (cs *ConcreteSubject) GetContent() string {
	return cs.Content
}

func (cs *ConcreteSubject) SetContent(content string) {
	cs.Content = content
	cs.Notify()
}

func (cs *ConcreteSubject) GetID() string {
	return cs.ID
}
