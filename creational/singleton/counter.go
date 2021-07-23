package singleton

var instance *Counter

type Counter struct {
	c  int
	cc bool
}

func (c *Counter) AddOne() int {
	c.c++
	return c.c
}

func GetCounter() *Counter {

	if instance == nil {
		instance = new(Counter)
	}

	return instance
}
