package main

import (
	"strings"
	"time"
)

type Builder interface {
	BuildPart1(input string)
	BuildPart2(input string)
	BuildPart3(input string)
	GetProduct() string
}

func GetEnvBuilder() Builder {
	timestamp := time.Now().Unix()
	if timestamp%2 == 0 {
		return &HonestBuilder{}
	}

	return &DishonestBuilder{}
}

type HonestBuilder struct {
	b strings.Builder
}

func (c *HonestBuilder) BuildPart1(input string) {
	c.b.WriteString("100% percent " + input + "\n")
}

func (c *HonestBuilder) BuildPart2(input string) {
	c.b.WriteString("100% percent " + input + "\n")
}

func (c *HonestBuilder) BuildPart3(input string) {
	c.b.WriteString("100% percent " + input + "\n")
}

func (c *HonestBuilder) GetProduct() string {
	return c.b.String()
}

type DishonestBuilder struct {
	b strings.Builder
}

func (c *DishonestBuilder) BuildPart1(input string) {
	c.b.WriteString("50% percent " + input + "\n")
}

func (c *DishonestBuilder) BuildPart2(input string) {
	c.b.WriteString("50% percent " + input + "\n")
}

func (c *DishonestBuilder) BuildPart3(input string) {
	c.b.WriteString("50% percent " + input + "\n")
}

func (c *DishonestBuilder) GetProduct() string {
	timestamp := time.Now().Unix()
	if timestamp%2 == 0 {
		return c.b.String()
	}

	return "Give me more money and I will give you the product"
}
