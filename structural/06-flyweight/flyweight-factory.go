package main

type FlyweightFactory struct {
	flyweights map[string]Flyweight
	shareable  map[string]bool
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
		shareable: map[string]bool{
			"A": true,
		},
	}
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if !f.shareable[key] {
		return &UnsharedConcreteFlyweight{State: key}
	}
	if _, ok := f.flyweights[key]; !ok {
		f.flyweights[key] = &ConcreteFlyweight{SharedState: key}
	}
	return f.flyweights[key]
}
