package simulator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Prototype(t *testing.T) {
	require.Equal(t, newSimulator(), Prototype())
}
