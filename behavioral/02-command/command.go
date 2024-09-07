package main

type Command interface {
	Execute()
	UnExecute()
}

var GlobalState = "init"

type ConcreteCommand struct {
	stateStack []string
	receiver   Receiver
}

func NewConcreteCommand(r Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		receiver:   r,
		stateStack: make([]string, 0),
	}
}

// Execute calls the receiver's Action method to do something, and additional functions.
// So the real work is done by the receiver, command just calls the receiver's method.
func (cc *ConcreteCommand) Execute() {
	cc.stateStack = append(cc.stateStack, GlobalState)
	GlobalState = cc.receiver.Action()
}

func (cc *ConcreteCommand) UnExecute() {
	GlobalState = cc.stateStack[len(cc.stateStack)-1]
	cc.stateStack = cc.stateStack[:len(cc.stateStack)-1]
}
