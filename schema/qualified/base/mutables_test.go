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

	testProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Data"))
	testProperties := baseLists.NewPropertyList(testProperty)
	testMutables := mutables{testProperties}
	require.Equal(t, mutables{PropertyList: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.GetMutablePropertyList())
	mutatedTestProperty := baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("Data2"))
	require.Equal(t, mutables{PropertyList: baseLists.NewPropertyList(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}
