package main

type IteratorType int

const (
	NormalIteratorType IteratorType = 1
	OddIteratorType    IteratorType = 2
)

type Aggregate interface {
	CreateIterator(t IteratorType) Iterator
}

func GiveMeAggregate() Aggregate {
	return &ItemsContainer{
		items: []*Item{
			{Name: "item1"},
			{Name: "item2"},
			{Name: "item3"},
			{Name: "item4"},
			{Name: "item5"},
			{Name: "item6"},
		},
	}
}

type Item struct {
	Name string
}

type ItemsContainer struct {
	items []*Item
}

func (i *ItemsContainer) CreateIterator(t IteratorType) Iterator {
	switch t {
	case NormalIteratorType:
		return NewNormalIterator(i.items)
	case OddIteratorType:
		return NewOddIterator(i.items)
	default:
		return nil
	}

}
