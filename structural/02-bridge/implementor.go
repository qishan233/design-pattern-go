package main

type Implementor interface {
	OperationImpl() string
}

type ConcreteImplementorA struct{}

func (c *ConcreteImplementorA) OperationImpl() string {
	return "ConcreteImplementorA OperationImpl"
}

type ConcreteImplementorB struct{}

func (c *ConcreteImplementorB) OperationImpl() string {
	return "ConcreteImplementorB OperationImpl"
}
