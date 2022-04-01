// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
)

func TestAuxiliary(t *testing.T) {
	context, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Auxiliary := NewAuxiliary("testAuxiliary", base.TestAuxiliaryKeeperPrototype).Initialize(Mapper, nil).(auxiliary)
	require.Equal(t, "testAuxiliary", Auxiliary.GetName())
	require.Equal(t, nil, Auxiliary.GetKeeper().Help(context, nil))
}
