package main

import "fmt"

type ScrollBar interface {
	Scroll()
}

type MotifScrollBar struct {
}

func (m *MotifScrollBar) Scroll() {
	fmt.Println("MotifScrollBar")
}

type PMScrollBar struct {
}

func (p *PMScrollBar) Scroll() {
	fmt.Println("PMScrollBar")
}
