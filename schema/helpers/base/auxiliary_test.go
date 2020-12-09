package base

import (
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"testing"
)

func TestAuxiliary(t *testing.T) {
	_, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Auxiliary := NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype)
	Auxiliary.Initialize(Mapper, nil)
	Auxiliary.GetName()
	Auxiliary.GetKeeper()
}
