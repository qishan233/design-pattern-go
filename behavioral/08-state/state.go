package main

import "fmt"

type State interface {
	Handle(ctx Context)
}

type StateA struct{}

func (s *StateA) Handle(ctx Context) {
	fmt.Println("StateA handle the request now, and change the state to StateB")
	ctx.SetState(&StateB{})
}

type StateB struct{}

func (s *StateB) Handle(ctx Context) {
	fmt.Println("StateB handle the request now, and change the state to StateC")
	ctx.SetState(&StateC{})
}

type StateC struct{}

func (s *StateC) Handle(ctx Context) {
	fmt.Println("StateC handle the request now, and change the state to StatA")
	ctx.SetState(&StateA{})
}
