package builder

type Person struct {
	Age     int
	Name    string
	Phone   string
	Address string
}

type PersonBuilder struct {
	age                  int
	name, phone, address string
}

func NewPersonBuilder(name string) *PersonBuilder {
	return &PersonBuilder{name: name}
}

func (b *PersonBuilder) SetAge(age int) *PersonBuilder {
	b.age = age
	return b
}

func (b *PersonBuilder) SetName(name string) *PersonBuilder {
	b.name = name
	return b
}

func (b *PersonBuilder) SetPhone(phone string) *PersonBuilder {
	b.phone = phone
	return b
}

func (b *PersonBuilder) SetAddress(address string) *PersonBuilder {
	b.address = address
	return b
}

func (b PersonBuilder) Build() Person {
	return Person{
		Age:     b.age,
		Name:    b.name,
		Phone:   b.phone,
		Address: b.address,
	}
}
