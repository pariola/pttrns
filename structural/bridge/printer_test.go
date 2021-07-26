package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBridge(t *testing.T) {

	// create printers
	hp := HPPrinter{}
	epson := EpsonPrinter{}

	var result string

	result = hp.Print("Hi!")
	assert.Equal(t, "HP Print: Hi!\n", result)

	result = epson.Print("Hey!")
	assert.Equal(t, "Epson Print: Hey!\n", result)

	// create computers
	mac := Mac{}
	windows := Windows{}

	mac.SetPrinter(hp)
	result = mac.Print("Mac Hi!")
	assert.Equal(t, "HP Print: Mac Hi!\n", result)

	mac.SetPrinter(epson)
	result = mac.Print("Mac Hi!")
	assert.Equal(t, "Epson Print: Mac Hi!\n", result)

	windows.SetPrinter(hp)
	result = windows.Print("Windows Hi!")
	assert.Equal(t, "HP Print: Windows Hi!\n", result)

	windows.SetPrinter(epson)
	result = windows.Print("Windows Hi!")
	assert.Equal(t, "Epson Print: Windows Hi!\n", result)
}
