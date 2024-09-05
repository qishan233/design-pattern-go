package main

type Abstraction interface {
	Operation() string
}

type RefinedAbstraction struct {
	Implementor Implementor
}

func NewRefinedAbstraction(implementor Implementor) *RefinedAbstraction {
	return &RefinedAbstraction{
		Implementor: implementor,
	}
}



func (r *RefinedAbstraction) Operation() string {
	return r.Implementor.OperationImpl()
}