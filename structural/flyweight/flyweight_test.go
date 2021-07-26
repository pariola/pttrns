package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlyweight(t *testing.T) {

	factory := TeamFactory{
		teams: make(map[TeamType]*Team),
	}

	one := factory.GetTeam(A)

	require.NotNil(t, one, "team can't be nil")

	two := factory.GetTeam(A)

	assert.Equal(t, one, two, "one & two must point to the same team")

	three := factory.GetTeam(B)

	require.NotNil(t, three, "team can't be nil")
}
