package auxiliaries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("supplement").GetName(), base.NewAuxiliaries(
		scrub.Auxiliary,
		supplement.Auxiliary,
	).Get("supplement").GetName())
}
