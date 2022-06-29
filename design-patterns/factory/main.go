package main

import "log"

func main() {
	// Main function does not care about the implementation of the circle and square. It only cares about creating those using the factory.

	//  One way to initialize
	circle := ShapeFactory("circle")
	square := ShapeFactory("square")

	// Another way to initialize
	var shape3 Shape = ShapeFactory("circle")

	circle.Area()
	square.Area()
	shape3.Area()
}

//  Define the abstract interface
type Shape interface {
	Area()
}

// Define the concrete class (struct). It implements the shape interface.
type Circle struct {
}

func (c *Circle) Area() {
	log.Println("Area of circle is π*r*r")
}

// Define the concrete class. It implements the shape interface.
type Square struct {
}

func (s *Square) Area() {
	log.Println("Area of square is r*r")
}

// Define the shape factory
func ShapeFactory(shape string) Shape {
	switch shape {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	}
	return nil
}
