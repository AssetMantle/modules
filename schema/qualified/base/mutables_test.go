// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Mutables(t *testing.T) {

	testProperty := baseTypes.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testProperties := base.NewPropertyList(testProperty)
	testMutables := Mutables{testProperties}
	require.Equal(t, Mutables{Properties: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.GetMutablePropertyList())
	mutatedTestProperty := baseTypes.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data2"))
	require.Equal(t, Mutables{Properties: base.NewPropertyList(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}
