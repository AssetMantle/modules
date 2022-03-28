package key

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype(), assetIDFromInterface(base.NewID("")))
}
