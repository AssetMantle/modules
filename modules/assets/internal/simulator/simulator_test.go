package simulator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_newSimulator(t *testing.T) {
	require.Equal(t, newSimulator(), simulator{})
}
