package main

// Template only define the skeleton of an algorithm. Does not implement the steps.This is a little different from others traditional oop languages.
// TemplateMethod is a method that defines the steps of an algorithm and allows subclasses to provide the implementation for one or more steps.
type Template interface {
	TemplateMethod()
}

// PrimaryOperation is an interface that defines the steps of an algorithm. Subclasses can implement these steps.
type PrimaryOperation interface {
	Step1()
	Step2()
	Step3()
}

type BaseTemplate struct {
	p PrimaryOperation
}

func (b *BaseTemplate) TemplateMethod() {
	println("BaseTemplate TemplateMethod has do something, now call the step 1")
	b.p.Step1()

	println("BaseTemplate TemplateMethod has do something, now call the step 2")
	b.p.Step2()

	println("BaseTemplate TemplateMethod has do something, now call the step 2")
	b.p.Step3()
}
