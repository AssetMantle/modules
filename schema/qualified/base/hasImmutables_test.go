// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/AssetMantle/modules/schema/types/base"

	"github.com/stretchr/testify/require"

	metaUtilities "github.com/AssetMantle/modules/utilities/meta"
)

func Test_HasImmutables(t *testing.T) {
	testProperty := base.NewProperty(base.NewID("ID"), base.NewHeightData(base.NewHeight(123)))
	testImmutables := HasImmutables{base.NewProperties(testProperty)}

	require.Equal(t, HasImmutables{Properties: base.NewProperties(testProperty)}, testImmutables)
	require.Equal(t, base.NewProperties(testProperty), testImmutables.GetImmutableProperties())
	require.Equal(t, base.NewID(metaUtilities.Hash([]string{testProperty.GetHashID().String()}...)), testImmutables.GenerateHashID())
	require.Equal(t, base.NewID(""), HasImmutables{base.NewProperties()}.GenerateHashID())
}
