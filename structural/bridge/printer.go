package bridge

import (
	"fmt"
)

type Printer interface {
	Print(string) string
}

type HPPrinter struct{}

func (p HPPrinter) Print(data string) string {
	return fmt.Sprintf("HP Print: %s\n", data)
}

type EpsonPrinter struct{}

func (p EpsonPrinter) Print(data string) string {
	return fmt.Sprintf("Epson Print: %s\n", data)
}

type Computer interface {
	Print(string) string
	SetPrinter(Printer)
}

type Mac struct {
	p Printer
}

func (m *Mac) SetPrinter(p Printer) {
	m.p = p
}

func (m Mac) Print(data string) string {
	return m.p.Print(data)
}

type Windows struct {
	p Printer
}

func (w *Windows) SetPrinter(p Printer) {
	w.p = p
}

func (w Windows) Print(data string) string {
	return w.p.Print(data)
}
