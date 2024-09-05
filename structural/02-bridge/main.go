package main

func main() {
	// we can use different implementor in different scene, and the Operation function is independent of the implementor
	abstractionA := NewRefinedAbstraction(&ConcreteImplementorA{})
	Operation("A implementor: ", abstractionA)

	abstractionB := NewRefinedAbstraction(&ConcreteImplementorB{})
	Operation("B implementor", abstractionB)
}

// Operation is independent of the implementor, only knows the abstraction. We can change implementor without changing the Operation function
func Operation(tips string, r Abstraction) {
	println(tips, r.Operation())
}
