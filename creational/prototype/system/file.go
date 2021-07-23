package system

import (
	"fmt"
)

type file struct {
	name string
}

func (f file) print(inden string) {
	fmt.Printf("%s%s\n", inden, f.name)
}

func (f file) clone() node {
	f.name = f.name + "_copy"
	return f
}
