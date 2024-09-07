package main

import "fmt"

func main() {
	factory := NewFlyweightFactory()

	// flyweightA is shareable
	flyweightA := factory.GetFlyweight("A")
	flyweightA1 := factory.GetFlyweight("A")

	fmt.Println("flyweightA and flyweightA1 are the same instance: ", flyweightA == flyweightA1)

	// flyweightB and flyweightC are not shareable
	flyweightB := factory.GetFlyweight("B")
	flyweightC := factory.GetFlyweight("C")

	fmt.Println("flyweightB and flyweightC are not the same instance: ", flyweightB != flyweightC)

	fmt.Println(flyweightA.Operation("A"))
	fmt.Println(flyweightA1.Operation("A1"))

	fmt.Println(flyweightB.Operation("B"))
	fmt.Println(flyweightB.Operation("B"))

	fmt.Println(flyweightC.Operation("C"))

}
