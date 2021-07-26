package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	maker := NewShapeMaker()

	maker.drawCircle()
	maker.drawSquare()
	maker.drawRectangle()
}
