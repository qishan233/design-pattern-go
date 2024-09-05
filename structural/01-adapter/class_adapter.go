package main

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type ClassAdapter struct{}

func NewClassAdapter() *ClassAdapter {
	return &ClassAdapter{}
}

func (a *ClassAdapter) Request() string {
	return a.SpecificRequest() + ", class adapter do nothing"
}

func (a *ClassAdapter) SpecificRequest() string {
	return "SpecificRequest"
}
