package factory

import (
	"fmt"
)

type Shape interface {
	draw()
}

type ShapeType int

const (
	Circle ShapeType = iota
	Square
	Rectangle
)

type circle struct{}
type square struct{}
type rectangle struct{}

// Draw circle
func (c circle) draw() {
	fmt.Println("circle:drawing")
}

// Draw square
func (s square) draw() {
	fmt.Println("square:drawing")
}

// Draw rectangle
func (r rectangle) draw() {
	fmt.Println("rectangle:drawing")
}

// NewShape is a interface factory design that supports multiple implementations through the use of interfaces.
func NewShape(t ShapeType) Shape {

	switch t {
	case Circle:
		return &circle{}
	case Square:
		return &square{}
	case Rectangle:
		return &rectangle{}
	}

	return nil
}
