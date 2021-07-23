package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinter(t *testing.T) {

	msg := "hello world!"
	client := &Client{}

	adapter := &LegacyPrinterAdapter{
		stored: msg,
		p:      &OneLegacyPrinter{},
	}

	printed := client.Print(adapter)
	assert.Equal(t, "Legacy Printer: hello world!\n", printed)
}
