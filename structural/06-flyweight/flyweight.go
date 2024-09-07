package main

type Flyweight interface {
	Operation(key string) string
	Shareable() bool
}

type ConcreteFlyweight struct {
	SharedState string
}

func (cf *ConcreteFlyweight) Operation(key string) string {
	return key + " Concrete Flyweight " + cf.SharedState
}
func (cf *ConcreteFlyweight) Shareable() bool {
	return true
}

type UnsharedConcreteFlyweight struct {
	State string
}

func (uf *UnsharedConcreteFlyweight) Operation(key string) string {
	uf.State += " More Money"
	return key + " Unshared Flyweight" + uf.State
}

func (uf *UnsharedConcreteFlyweight) Shareable() bool {
	return false
}
