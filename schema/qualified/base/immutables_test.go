// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	metaUtilities "github.com/AssetMantle/modules/utilities/meta"
)

func Test_Immutables(t *testing.T) {
	testProperty := baseTypes.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
	testImmutables := Immutables{baseTypes.NewProperties(testProperty)}

	require.Equal(t, Immutables{Properties: baseTypes.NewProperties(testProperty)}, testImmutables)
	require.Equal(t, baseTypes.NewProperties(testProperty), testImmutables.GetImmutableProperties())
	require.Equal(t, baseIDs.NewID(metaUtilities.Hash([]string{testProperty.GetHash().String()}...)), testImmutables.GenerateHashID())
	require.Equal(t, baseIDs.NewID(""), Immutables{baseTypes.NewProperties()}.GenerateHashID())
}
