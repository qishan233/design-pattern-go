package main

import "fmt"

type Director struct{}

func (d *Director) Construct() {
	// this is Factory Method pattern
	builder := GetEnvBuilder()

	// the same construction process can build different products, if the builder is different, even though the result is the same: a string
	builder.BuildPart1("mountain")
	builder.BuildPart2("river")
	builder.BuildPart3("tree")

	fmt.Println(builder.GetProduct())
}
