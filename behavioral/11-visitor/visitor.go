package main

import (
	"fmt"
)

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitorA struct {
}

func (c *ConcreteVisitorA) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("ConcreteVisitorA visit ConcreteElementA, name:%s\n", element.Name)
}

func (c *ConcreteVisitorA) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("ConcreteVisitorA visit ConcreteElementB, age:%d\n", element.Age)
}

type ConcreteVisitorB struct {
}

func (c *ConcreteVisitorB) VisitConcreteElementA(element *ConcreteElementA) {

	fmt.Printf("ConcreteVisitorB visit ConcreteElementA, name:%s-M\n", element.Name)
}

func (c *ConcreteVisitorB) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("ConcreteVisitorB visit ConcreteElementB, age:%d\n", element.Age+10)
}
