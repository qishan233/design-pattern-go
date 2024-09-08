package main

import (
	"math/rand"
)

type Context struct {
}

func (c *Context) DoSomething() {
	var strategy Strategy

	if rand.Intn(10)%2 == 0 {
		strategy = &HappyStrategy{}
	} else {
		strategy = &SadStrategy{}
	}

	println("prepare step one")
	strategy.Step1()

	println("prepare step two")
	strategy.Step2()

	println("done")
}
