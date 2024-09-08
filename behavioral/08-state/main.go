package main

func main() {
	stateB := &StateB{}

	ctx := NewStateContext()

	ctx.SetState(stateB)

	for i := 0; i < 5; i++ {
		ctx.Request()
	}
}
