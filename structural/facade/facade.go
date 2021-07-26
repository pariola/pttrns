package facade

import (
	"fmt"
)

type Shape interface {
	draw()
}

type Circle struct{}
type Square struct{}
type Rectangle struct{}

func (Circle) draw() {
	fmt.Println("draw:circle")
}

func (Square) draw() {
	fmt.Println("draw:square")
}

func (Rectangle) draw() {
	fmt.Println("draw:rectangle")
}

type ShapeMaker struct {
	c Circle
	s Square
	r Rectangle
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		c: Circle{},
		s: Square{},
		r: Rectangle{},
	}
}

func (m ShapeMaker) drawCircle() {
	m.c.draw()
}

func (m ShapeMaker) drawSquare() {
	m.s.draw()
}

func (m ShapeMaker) drawRectangle() {
	m.r.draw()
}
