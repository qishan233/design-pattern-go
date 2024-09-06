package main

func main() {
	// Create a tree structure
	root := &Composite{Name: "root"}
	la := &Leaf{Name: "Leaf-A"}
	lb := &Leaf{Name: "Leaf-B"}
	_ = root.Add(la)
	_ = root.Add(lb)

	branchA := &Composite{Name: "Branch A"}
	lc := &Leaf{Name: "Leaf-C"}
	ld := &Leaf{Name: "Leaf-D"}
	_ = branchA.Add(lc)
	_ = branchA.Add(ld)

	_ = root.Add(branchA)

	branchB := &Composite{Name: "Branch B"}
	le := &Leaf{Name: "Leaf-E"}
	lf := &Leaf{Name: "Leaf-F"}
	_ = branchB.Add(le)
	_ = branchB.Add(lf)

	_ = root.Add(branchB)

	Operate(root)

}

// Operate is a client function that operates on a component, no different between a leaf or a composite.
// In another words, the client does not need to know the type of the component.
func Operate(c Component) {
	c.Operation("")
}
