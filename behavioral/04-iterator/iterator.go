package main

type Iterator interface {
	Next() *Item
	HasNext() bool
	Current() *Item
}

type NormalIterator struct {
	items  []*Item
	cursor int
}

func NewNormalIterator(items []*Item) *NormalIterator {
	return &NormalIterator{
		items:  items,
		cursor: -1,
	}
}

func (n *NormalIterator) Next() *Item {
	if !n.HasNext() {
		return nil
	}

	n.cursor++

	return n.items[n.cursor]
}

func (n *NormalIterator) HasNext() bool {
	return n.cursor+1 < len(n.items)
}

func (n *NormalIterator) Current() *Item {
	if n.cursor < 0 {
		// iterator not started,cursor point to nil
		return nil
	}
	return n.items[n.cursor]
}

type OddIterator struct {
	items  []*Item
	cursor int
}

func NewOddIterator(items []*Item) *OddIterator {
	return &OddIterator{
		items:  items,
		cursor: -2,
	}
}

func (o *OddIterator) Next() *Item {
	if !o.HasNext() {
		return nil
	}

	o.cursor += 2

	return o.items[o.cursor]
}

func (o *OddIterator) HasNext() bool {
	return o.cursor+2 < len(o.items)
}

func (o *OddIterator) Current() *Item {
	if o.cursor < 0 {
		// iterator not started,cursor point to nil
		return nil
	}
	return o.items[o.cursor]
}
