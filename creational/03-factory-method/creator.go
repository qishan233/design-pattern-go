package main

type Creator interface {
	CreateProduct(money int) Product
}

type MagicalCreator struct{}

func (c *MagicalCreator) CreateProduct(money int) Product {
	if money > 5000 {
		return NewProductHighQuality(money)
	}

	return NewProductLowQuality(money)
}
