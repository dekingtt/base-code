package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, long float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.long
}

func (r rectangle) perim() float64 {
	return r.width*2 + r.long*2
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectangle{width: 2.5, long: 3.8}
	c := circle{radius: 8.88}

	measure(r)
	measure(c)
}
