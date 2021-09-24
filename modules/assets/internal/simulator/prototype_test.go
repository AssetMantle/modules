package simulator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Prototype(t *testing.T) {
	require.Equal(t, newSimulator(), Prototype())
}
