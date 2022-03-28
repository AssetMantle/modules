package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterCodec(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, RegisterLegacyAminoCodec(nil), nil)
	})
}
