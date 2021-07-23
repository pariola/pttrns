package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerson(t *testing.T) {

	p := NewPerson("Blessing", 21)

	assert.Equal(t, 21, p.age, "person age should match")
	assert.Equal(t, "Blessing", p.name, "person name should match")
}
