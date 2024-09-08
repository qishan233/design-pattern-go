package main

type PrimaryOperationImpl struct {
}

func NewTemplate() Template {
	p := &PrimaryOperationImpl{}

	t := &BaseTemplate{
		p: p,
	}

	return t
}

func (p *PrimaryOperationImpl) Step1() {
	println("PrimaryStepImpl Step1")
}

func (p *PrimaryOperationImpl) Step2() {
	println("PrimaryStepImpl Step2")
}
func (p *PrimaryOperationImpl) Step3() {
	println("PrimaryStepImpl Step3")
}
