package main

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct {
	Name string
}

// Accept method is used to call the visitor's visit method. The logic of how to use info is encapsulated in the visitor, not in the element.
func (c *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(c)
}

type ConcreteElementB struct {
	Age int
}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(c)
}
