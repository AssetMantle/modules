package simulator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_newSimulator(t *testing.T) {
	require.Equal(t, newSimulator(), simulator{})
}
