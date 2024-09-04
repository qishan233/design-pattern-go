package main

import "time"

type Product interface {
	GetPrice() int
	GetQuality() string
}

type ProductHighQuality struct {
	price int
}

func NewProductHighQuality(price int) *ProductHighQuality {
	return &ProductHighQuality{price: price}
}

func (p *ProductHighQuality) GetPrice() int {
	return p.price
}

func (p *ProductHighQuality) GetQuality() string {
	return "High Quality"
}

type ProductLowQuality struct {
	realPrice int
}

func NewProductLowQuality(price int) *ProductLowQuality {
	return &ProductLowQuality{
		realPrice: price,
	}
}

func (p *ProductLowQuality) GetPrice() int {
	return p.realPrice / 10
}

func (p *ProductLowQuality) GetQuality() string {
	if time.Now().Unix()%10 == 0 {
		return "High Quality"
	}

	return "Low Quality"
}
