package factory

type Person struct {
	name string
	age  int
}

// NewPerson is a simple factory design
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}
