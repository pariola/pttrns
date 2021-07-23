package adapter

import (
	"fmt"
)

type Mac struct{}

func (m Mac) InsertSquareUSB() {
	fmt.Println("insert square USB into mac!")
}

type Windows struct{}

func (w Windows) InsertCircleUSB() {
	fmt.Println("insert circle USB into windows!")
}

type WindowsSquareAdapter struct {
	W *Windows
}

func (a WindowsSquareAdapter) InsertSquareUSB() {
	a.W.InsertCircleUSB()
}
