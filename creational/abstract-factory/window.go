package main

import "fmt"

type Window interface {
	Draw()
}

type PMWindow struct{}

func (p *PMWindow) Draw() {
	fmt.Println("PMWindow Draw")
}

type MotifWindow struct{}

func (m *MotifWindow) Draw() {
	fmt.Println("MotifWindow Draw")
}
