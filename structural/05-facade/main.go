package main

func main() {
	robot := &Root{}
	
	// client code doesn't know how to make dinner, all steps are hidden inside the robot, thus robot is a facade
	robot.MakeDinner()
}
