package prototype

type Person struct {
	name, email string
}

func NewPerson(name, email string) *Person {
	return &Person{name: name, email: email}
}

// WithName clones the specified `Person` struct and updates the name value.
func (p Person) WithName(name string) Person {
	p.name = name
	return p
}

// WithEmail clones the specified `Person` struct and updates the email value.
func (p Person) WithEmail(email string) Person {
	p.email = email
	return p
}
