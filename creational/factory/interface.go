package factory

import (
	"fmt"
)

type animal struct {
	name string
}

func (a animal) Eat() {
	fmt.Println("i'm eating!")
}

type Animal interface {
	Eat()
}

// NewAnimal is an interface factory design
func NewAnimal(name string) Animal {
	return &animal{name: name}
}
