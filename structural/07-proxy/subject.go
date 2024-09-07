package main

import "fmt"

type Subject interface {
	Request()
}

type RealSubject struct {
}

func (rs *RealSubject) Request() {
	fmt.Println("RealSubject handle request")
}
