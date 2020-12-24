package base

import (
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuxiliary(t *testing.T) {
	context, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Auxiliary := NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil).(auxiliary)
	require.Equal(t, "testAuxiliary", Auxiliary.GetName())
	require.Equal(t, nil, Auxiliary.GetKeeper().Help(context, nil))
}
