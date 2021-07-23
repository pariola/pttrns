package factory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewShape(t *testing.T) {
	require.IsType(t, &circle{}, NewShape(Circle), "shape should be type `circle`")
	require.IsType(t, &square{}, NewShape(Square), "shape should be type `square`")
	require.IsType(t, &rectangle{}, NewShape(Rectangle), "shape should be type `rectangle`")
}
