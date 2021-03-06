/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package qualified

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"

	"github.com/stretchr/testify/require"

	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

func Test_HasImmutables(t *testing.T) {
	testProperty := base.NewProperty(base.NewID("ID"), base.NewHeightData(base.NewHeight(123)))
	testImmutables := HasImmutables{base.NewProperties(testProperty)}

	require.Equal(t, HasImmutables{Properties: base.NewProperties(testProperty)}, testImmutables)
	require.Equal(t, base.NewProperties(testProperty), testImmutables.GetImmutableProperties())
	require.Equal(t, base.NewID(metaUtilities.Hash([]string{testProperty.GetHashID().String()}...)), testImmutables.GenerateHashID())
	require.Equal(t, base.NewID(""), HasImmutables{base.NewProperties()}.GenerateHashID())
}
