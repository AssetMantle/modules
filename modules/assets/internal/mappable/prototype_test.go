package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype(), asset{})
}
