// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Mutables(t *testing.T) {

	testProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testProperties := baseLists.NewPropertyList(testProperty)
	testMutables := Mutables{testProperties}
	require.Equal(t, Mutables{Properties: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.GetMutablePropertyList())
	mutatedTestProperty := baseProperties.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data2"))
	require.Equal(t, Mutables{Properties: baseLists.NewPropertyList(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}
