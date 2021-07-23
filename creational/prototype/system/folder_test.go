package system

import (
	"testing"
)

func TestSystem(t *testing.T) {
	f1 := &file{name: "File1"}
	f2 := &file{name: "File2"}
	f3 := &file{name: "File3"}

	fold1 := &folder{
		name:  "Folder1",
		items: []node{f1},
	}

	fold2 := &folder{
		name:  "Folder2",
		items: []node{fold1, f2, f3},
	}

	fold2.print("  ")

	// clone
	cloned := fold2.clone()
	cloned.print("  ")
}
