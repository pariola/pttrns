package system

import (
	"fmt"
)

type folder struct {
	name  string
	items []node
}

func (f folder) print(inden string) {
	fmt.Printf("%s%s\n", inden, f.name)
	for _, i := range f.items {
		i.print(inden + inden)
	}
}

func (f folder) clone() node {
	f.name = f.name + "_copy"

	for i, item := range f.items {
		f.items[i] = item.clone()
	}

	return f
}
