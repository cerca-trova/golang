package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func main() {
	var c Circle      // define a struct
	c1 := new(Circle) // use keynew new for definition

	// Initialization
	c = Circle{x: 0, y: 0, r: 5}

	fmt.Println(c.r, c.x, c.y)
	fmt.Println(c1.r, c1.x, c1.y)
	fmt.Println(circleArea(c))
	fmt.Println(c.area())
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Interface
type totalArea interface {
	area() float64
}

// Interface  implementation
func totalAreaInFunc(shapes ...totalArea) float64 {
	var area float64
	for _, v := range shapes {
		area += v.area()
	}
	return area
}
