package main

type ObjectStructure struct {
	elements []Element
}

// Visit method is used to call the visitor's visit method.
// The logic of how to visit elements'info is encapsulated in the visitor.
// Visitor does not need to know the internal structure of this object.
func (o *ObjectStructure) Visit(visitor Visitor) {
	for _, element := range o.elements {
		element.Accept(visitor)
	}
}
