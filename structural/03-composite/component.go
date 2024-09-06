package main

type Component interface {
	Operation(prefix string)
	Add(c Component) error
	Remove(c Component) error
	GetChild(i int) Component
}
