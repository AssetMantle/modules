/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	moduleConstants "github.com/persistenceOne/persistenceSDK/constants/modules"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
)

func Test_DecID(t *testing.T) {
	testID := NewDecID(sdkTypes.OneDec().Neg())

	require.Equal(t, decID{Value: sdkTypes.OneDec().Neg()}, testID)
	require.Equal(t, sdkTypes.OneDec().Neg().String(), testID.String())
	require.Equal(t, true, testID.Equals(testID))
	require.Equal(t, false, testID.Equals(NewID("ID2")))
	require.Equal(t, append([]byte{moduleConstants.NegativeExchangeRate, uint8(1)}, []byte("1.000000000000000000")...), testID.Bytes())

	testID = NewDecID(sdkTypes.OneDec())
	require.Equal(t, append([]byte{moduleConstants.PositiveExchangeRate, uint8(1)}, []byte("1.000000000000000000")...), testID.Bytes())
}
