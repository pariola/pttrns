package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {

	one := NewPerson("ABC", "abc@gmail.com")

	assert.Equal(t, "ABC", one.name, "Person name don't match")
	assert.Equal(t, "abc@gmail.com", one.email, "Person email don't match")

	// clone WithName
	two := one.WithName("XYZ")

	assert.Equal(t, "abc@gmail.com", two.email, "Cloned Person email don't match")
	assert.Equal(t, "XYZ", two.name, "Cloned Person name doesn't match")
	assert.Equal(t, "ABC", one.name, "Original Person name doesn't match")

	// clone WithEmail
	three := one.WithEmail("abc2@gmail.com")

	assert.Equal(t, "ABC", three.name, "Cloned Person name don't match")
	assert.Equal(t, "abc2@gmail.com", three.email, "Cloned Person email doesn't match")
	assert.Equal(t, "abc@gmail.com", one.email, "Original Person email doesn't match")
}
