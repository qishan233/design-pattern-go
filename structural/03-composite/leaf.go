package main

import "errors"

type Leaf struct {
	Name string
}

func (l *Leaf) Operation(prefix string) {
	println(prefix + l.Name + "-Leaf operation")
}

func (l *Leaf) Add(c Component) error {
	return errors.New("Cannot add component to leaf")
}

func (l *Leaf) Remove(c Component) error {
	return errors.New("Cannot remove component from leaf")
}

func (l *Leaf) GetChild(i int) Component {
	return nil
}
