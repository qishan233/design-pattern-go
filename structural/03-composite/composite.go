package main

import "errors"

type Composite struct {
	Name     string
	children []Component
}

func (c *Composite) Operation(prefix string) {
	println(prefix + c.Name + "-Composite operation")
	for _, child := range c.children {
		child.Operation("	" + prefix)
	}
}

func (c *Composite) Add(child Component) error {
	c.children = append(c.children, child)
	return nil
}

func (c *Composite) Remove(child Component) error {
	for i, candidate := range c.children {
		if candidate == child {
			c.children = append(c.children[:i], c.children[i+1:]...)
			return nil
		}
	}
	return errors.New("Child not found")
}

func (c *Composite) GetChild(i int) Component {
	if i < 0 || i >= len(c.children) {
		return nil
	}
	return c.children[i]
}
