package main

func main() {
	o := &ObjectStructure{
		elements: []Element{
			&ConcreteElementA{Name: "elementA"},
			&ConcreteElementB{Age: 1},
		},
	}

	// same object structure can be visited by different visitors
	v1 := &ConcreteVisitorA{}
	o.Visit(v1)

	
	v2 := &ConcreteVisitorB{}
	o.Visit(v2)
}
