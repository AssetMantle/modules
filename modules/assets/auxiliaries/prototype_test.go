package auxiliaries

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Prototype(t *testing.T) {
	require.Equal(t, base.NewAuxiliaries(), Prototype())
}
