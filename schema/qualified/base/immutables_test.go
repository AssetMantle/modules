// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func Test_Immutables(t *testing.T) {
	testProperty := base2.NewProperty(baseIDs.NewStringID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
	testImmutables := Immutables{base.NewPropertyList(testProperty)}

	require.Equal(t, Immutables{PropertyList: base.NewPropertyList(testProperty)}, testImmutables)
	require.Equal(t, base.NewPropertyList(testProperty), testImmutables.GetImmutablePropertyList())
	require.Equal(t, baseIDs.NewStringID(stringUtilities.Hash([]string{testProperty.GetHash().String()}...)), testImmutables.GenerateHashID())
	require.Equal(t, baseIDs.NewStringID(""), Immutables{base.NewPropertyList()}.GenerateHashID())
}
