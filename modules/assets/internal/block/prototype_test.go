package block

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, block{}, Prototype())
}
