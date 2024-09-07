package main

import (
	"fmt"
	"math/rand"
)

type EvenProxy struct {
	s Subject
}

func NewEvenProxy(s Subject) *EvenProxy {
	return &EvenProxy{s: s}
}

func (ep *EvenProxy) Request() {
	r := rand.Intn(10)

	if r%2 == 0 && ep.s != nil {
		ep.s.Request()
		return
	}
	fmt.Println("subject is tired, please call this method later——EvenProxy")
}

type OddProxy struct {
	s Subject
}

func NewOddProxy(s Subject) *OddProxy {
	return &OddProxy{s: s}
}

func (op *OddProxy) Request() {
	r := rand.Intn(10)

	if r%2 == 0 && op.s != nil {
		op.s.Request()
		return
	}
	fmt.Println("subject is tired, please call this method later——OddProxy")
}
