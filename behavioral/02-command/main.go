package main

import "fmt"

func main() {
	invoker := NewInvoker()

	receiverA := &ConcreteReceiver{}

	// create a command
	command := &ConcreteCommand{
		receiver: receiverA,
	}

	commandId := "ChangeState"

	// set the command to the invoker
	invoker.SetCommand(commandId, command)

	fmt.Println("Current Global State: ", GlobalState)

	// execute the command
	invoker.ExecuteCommand(commandId)
	fmt.Println("ExecuteCommand, current Global State: ", GlobalState)

	invoker.ExecuteCommand(commandId)
	fmt.Println("ExecuteCommand, current Global State: ", GlobalState)

	invoker.UnExecuteCommand(commandId)
	fmt.Println("UnExecuteCommand, current Global State: ", GlobalState)

	invoker.UnExecuteCommand(commandId)
	fmt.Println("UnExecuteCommand, current Global State: ", GlobalState)
}
