package main

type Context interface {
	SetState(State)
	Request()
}

type StateContext struct {
	state State
}

func NewStateContext() Context {
	return &StateContext{}
}

func (sc *StateContext) SetState(s State) {
	sc.state = s
}

func (sc *StateContext) Request() {
	if sc.state == nil {
		sc.SetState(&StateA{})
	}

	sc.state.Handle(sc)
}
