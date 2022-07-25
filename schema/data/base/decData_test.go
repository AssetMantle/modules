// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/data/constants"
)

func Test_DecData(t *testing.T) {

	decValue := sdkTypes.NewDec(12)
	testDecData := NewDecData(decValue)

	require.Equal(t, decValue.String(), testDecData.String())
	require.Equal(t, constants.DecDataID, testDecData.GetType())
	require.Equal(t, false, testDecData.Compare(NewStringData("")) == 0)
	require.Equal(t, true, testDecData.Compare(NewDecData(sdkTypes.NewDec(12))) == 0)
}
