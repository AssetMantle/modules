package auxiliaries

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("supplement").GetName(), base.NewAuxiliaries(
		scrub.Auxiliary,
		supplement.Auxiliary,
	).Get("supplement").GetName())
}
