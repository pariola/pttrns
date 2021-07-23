package singleton

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCounter(t *testing.T) {

	c := GetCounter()

	require.NotNil(t, c, "Initialized Counter can't be nil")

	if c.AddOne() != 1 {
		t.Fatal("Counter `count` should be 1")
	}
}
