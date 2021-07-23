package adapter

import (
	"fmt"
)

// LegacyPrinter - old interface
type LegacyPrinter interface {
	Print(string) string
}

type OneLegacyPrinter struct{}

func (p OneLegacyPrinter) Print(s string) string {
	o := fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Print(o)
	return o
}

// ModernPrinter - new interface
type ModernPrinter interface {
	PrintCache() string
}

type OneModernPrinter struct {
	stored string
}

func (p OneModernPrinter) PrintCache() string {
	o := fmt.Sprintf("Modern Printer: %s\n", p.stored)
	fmt.Print(o)
	return o
}

// Client
type Client struct{}

func (c Client) Print(p ModernPrinter) string {
	return p.PrintCache()
}

type LegacyPrinterAdapter struct {
	stored string
	p      LegacyPrinter
}

func (a LegacyPrinterAdapter) PrintCache() string {
	return a.p.Print(a.stored)
}
