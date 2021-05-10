package auxiliaries

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("verify").GetName(), base.NewAuxiliaries(
		verify.Auxiliary,
	).Get("verify").GetName())
}
