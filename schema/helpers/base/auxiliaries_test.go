package base

import (
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"testing"
)

func TestAuxiliaries(t *testing.T) {
	_, storeKey := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Auxiliaries := NewAuxiliaries(NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil))
	Auxiliaries.Get("testAuxiliary")
	Auxiliaries.GetList()
}
