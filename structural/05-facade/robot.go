package main

import "fmt"

type IRoot interface {
	MakeDinner() string
}

type Root struct{}

func (r *Root) MakeDinner() string {
	var (
		buyer  Buyer  = &buyer{}
		washer Washer = &washer{}
		cutter Cutter = &cutter{}
		cooker Cooker = &cooker{}
	)

	vegetables := buyer.BuyVegetables()
	washer.WashVegetables(vegetables)
	cutter.CutVegetables(vegetables)
	cooker.Cook(vegetables)

	fmt.Println("Dinner is ready!")

	return "Do it yourself!"
}

type Buyer interface {
	BuyVegetables() string
}

type buyer struct{}

func (b *buyer) BuyVegetables() string {
	fmt.Println("Buy vegetables: tomato, potato, onion")

	return "tomato, potato, onion"
}

type Washer interface {
	WashVegetables(string)
}

type washer struct{}

func (w *washer) WashVegetables(vegetables string) {
	fmt.Println("Wash " + vegetables)
}

type Cutter interface {
	CutVegetables(string)
}
type cutter struct{}

func (c *cutter) CutVegetables(vegetables string) {
	fmt.Println("Cut " + vegetables)
}

type Cooker interface {
	Cook(string)
}
type cooker struct{}

func (c *cooker) Cook(vegetables string) {
	fmt.Println("Cook " + vegetables)
}
