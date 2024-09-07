package main

type Invoker struct {
	commands map[string]Command
}

func NewInvoker() *Invoker {
	return &Invoker{
		commands: make(map[string]Command),
	}
}

func (i *Invoker) SetCommand(name string, c Command) {
	i.commands[name] = c
}

func (i *Invoker) ExecuteCommand(name string) {
	i.commands[name].Execute()
}

func (i *Invoker) UnExecuteCommand(name string) {
	i.commands[name].UnExecute()
}
