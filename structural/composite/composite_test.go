package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {

	goat := animal{}
	goat.eat() // eating!

	bolt := athlete{}
	bolt.train() // training!

	swimImpl := swimmerImpl{}

	hammerhead := shark{
		swimmer: swimmerImpl{},
	}
	hammerhead.eat()  // eating!
	hammerhead.swim() // swimming!

	phelps := swimmerAthlete{
		swimmer: swimImpl,
	}
	phelps.train() // training!
	phelps.swim()  // swimming!
}
