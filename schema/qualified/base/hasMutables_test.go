// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Mutables(t *testing.T) {

	testProperty := baseTypes.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data"))
	testProperties := baseTypes.NewProperties(testProperty)
	testMutables := HasMutables{testProperties}
	require.Equal(t, HasMutables{Properties: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.GetMutableProperties())
	mutatedTestProperty := baseTypes.NewProperty(baseIDs.NewID("ID"), baseData.NewStringData("Data2"))
	require.Equal(t, HasMutables{Properties: baseTypes.NewProperties(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}
