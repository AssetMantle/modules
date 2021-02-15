/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
)

func Test_Immutables(t *testing.T) {
	testProperty := NewProperty(NewID("ID"), NewFact(NewHeightData(NewHeight(123))))
	testImmutables := NewImmutables(NewProperties(testProperty))

	require.Equal(t, immutables{Properties: NewProperties(testProperty)}, testImmutables)
	require.Equal(t, NewProperties(testProperty), testImmutables.Get())
	require.Equal(t, id{IDString: metaUtilities.Hash([]string{testProperty.GetFact().GetHashID().String()}...)}, testImmutables.GenerateHashID())
	require.Equal(t, NewID(""), NewImmutables(NewProperties()).GenerateHashID())
}
