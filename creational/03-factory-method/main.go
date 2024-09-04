package main

func main() {
	creator := &MagicalCreator{}
	product := creator.CreateProduct(200)
	// product := creator.CreateProduct(2000)

	println(product.GetPrice())
	println(product.GetQuality())
}
