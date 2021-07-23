package builder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pttrns/creational/builder"
)

func TestPersonBuilder(t *testing.T) {

	b := builder.NewPersonBuilder("Blessing").
		SetAge(21).
		SetPhone("+234808693****").
		SetAddress("Somewhere in Lagos!")

	person := b.Build()

	assert.Equal(t, 21, person.Age, "Person `age` doesn't match")
	assert.Equal(t, "Blessing", person.Name, "Person `name` doesn't match")
	assert.Equal(t, "+234808693****", person.Phone, "Person `phone` doesn't match")
	assert.Equal(t, "Somewhere in Lagos!", person.Address, "Person `address` doesn't match")

	person = b.SetName("Frank").Build()

	assert.Equal(t, "Frank", person.Name, "Person `name` doesn't match")
}
