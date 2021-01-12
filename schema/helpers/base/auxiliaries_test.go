/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuxiliaries(t *testing.T) {
	_, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	require.Equal(t, false, base.KeyPrototype().Matches(base.NewKey("ID")))
	require.Equal(t, false, base.KeyPrototype().IsPartial())
	Auxiliaries := NewAuxiliaries(NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil)).(auxiliaries)
	require.Equal(t, "testAuxiliary", Auxiliaries.Get("testAuxiliary").GetName())
	require.Nil(t, nil, Auxiliaries.Get(""))
	Auxiliaries.GetList()
	require.Equal(t, "testAuxiliary", Auxiliaries.GetList()[0].GetName())
}
