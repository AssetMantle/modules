package transaction

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegisterCodec(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, RegisterCodec(nil), nil)
	})
}
