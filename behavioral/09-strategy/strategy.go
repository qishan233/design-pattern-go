package main

type Strategy interface {
	Step1()
	Step2()
}

type HappyStrategy struct{}

func (HappyStrategy) Step1() {
	println("HappyStrategy Step1")
}
func (HappyStrategy) Step2() {
	println("HappyStrategy Step2")
}

type SadStrategy struct{}

func (SadStrategy) Step1() {
	println("SadStrategy Step1")
}
func (SadStrategy) Step2() {
	println("SadStrategy Step2")
}
