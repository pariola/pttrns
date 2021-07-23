package adapter_test

import (
	"testing"

	"pttrns/structural/adapter"
)

type Computer interface {
	InsertSquareUSB()
}

type Client struct{}

func (Client) InsertUSB(c Computer) {
	c.InsertSquareUSB()
}

func TestUSBAdapter(t *testing.T) {

	mac := new(adapter.Mac)
	windows := new(adapter.Windows)

	c := new(Client)
	c.InsertUSB(mac)

	// can't use Windows directly as it's not supported
	// c.InsertUSB(windows)

	// adaptee -> windows
	// target -> Computer

	a := adapter.WindowsSquareAdapter{W: windows}
	c.InsertUSB(a)
}
