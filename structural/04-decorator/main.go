package main

import "fmt"

func main() {
	c := &ConcreteComponent{}

	decoratorA := NewConcreteDecoratorA(c)
	
	decoratorB := NewConcreteDecoratorB(decoratorA)

	fmt.Println(decoratorA.Operation())

	fmt.Println(decoratorB.Operation())
}
