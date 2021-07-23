package composite

import (
	"fmt"
)

// athlete someone that can train
type athlete struct{}

func (a athlete) train() {
	fmt.Println("training!")
}

type swimmer interface {
	swim()
}

type swimmerImpl struct {
}

func (i swimmerImpl) swim() {
	fmt.Println("swimming!")
}

// swimmerAthlete an athlete that also swims
type swimmerAthlete struct {
	athlete
	swimmer
}

// animal that eats
type animal struct{}

func (a animal) eat() {
	fmt.Println("eating!")
}

// shark a type of animal that can swim
type shark struct {
	animal
	swimmer
}
