package auxiliaries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, base.NewAuxiliaries(
		conform.Auxiliary,
		define.Auxiliary,
	).Get(""), Prototype().Get(""))
}
