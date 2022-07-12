// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/data/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/string"
)

func Test_DecData(t *testing.T) {

	decValue := sdkTypes.NewDec(12)
	testDecData := NewDecData(decValue)
	testDecData2 := NewDecData(sdkTypes.NewDec(0))

	require.Equal(t, decValue.String(), testDecData.String())
	require.Equal(t, baseIDs.NewStringID(string.Hash(decValue.String())), testDecData.GenerateHash())
	require.Equal(t, baseIDs.NewStringID(""), testDecData2.GenerateHash())
	require.Equal(t, constants.DecDataID, testDecData.GetType())
	require.Equal(t, false, testDecData.Compare(NewStringData("")) == 0)
	require.Equal(t, true, testDecData.Compare(NewDecData(sdkTypes.NewDec(12))) == 0)
}
