package factory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAnimal(t *testing.T) {

	a := NewAnimal("GOAT")

	require.NotNil(t, a, "animal should not be nil")

	a.Eat()
}
