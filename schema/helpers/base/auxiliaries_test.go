// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestAuxiliaries(t *testing.T) {
	_, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	require.Equal(t, false, base.KeyPrototype().Equals(base.NewKey("ID")))
	require.Equal(t, false, base.KeyPrototype().IsPartial())
	Auxiliaries := NewAuxiliaries(NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil)).(auxiliaries)
	require.Equal(t, "testAuxiliary", Auxiliaries.Get("testAuxiliary").GetName())
	require.Nil(t, nil, Auxiliaries.Get(""))
	Auxiliaries.GetList()
	require.Equal(t, "testAuxiliary", Auxiliaries.GetList()[0].GetName())
}
