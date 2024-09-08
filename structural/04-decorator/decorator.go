package main

type Decorator interface {
	Component
}

func NewConcreteDecoratorA(component Component) *ConcreteDecoratorA {
	return &ConcreteDecoratorA{component: component}
}

func NewConcreteDecoratorB(component Component) *ConcreteDecoratorB {
	return &ConcreteDecoratorB{component: component}
}

type ConcreteDecoratorA struct {
	component Component
}

func (c *ConcreteDecoratorA) Operation() string {
	if c.component != nil {
		return "ConcreteDecoratorA: " + c.component.Operation()
	}
	return "Decorator can't operate without a component"
}

type ConcreteDecoratorB struct {
	component Component
}

func (c *ConcreteDecoratorB) Operation() string {
	if c.component != nil {
		return "ConcreteDecoratorB: " + c.component.Operation()
	}

	return "Decorator can't operate without a component"
}
