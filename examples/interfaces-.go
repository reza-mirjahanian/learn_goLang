package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type reac2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// implement an interface in Go, we just need to implement all the methods in the interface. Here we implement geometry on rects.
func (r reac2) area() float64 {
	return r.width * r.height
}
func (r reac2) perim() float64 {
	return 2*r.width + 2*r.height
}

//Unlike other languages like Java, you don’t need to explicitly specify that a type implements an interface using something like an implements keyword. You just implement all the methods declared in the interface and you’re done.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := reac2{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
