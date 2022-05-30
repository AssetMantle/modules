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
	require.Equal(t, baseIDs.NewID(string.Hash(decValue.String())), testDecData.GenerateHash())
	require.Equal(t, baseIDs.NewID(""), testDecData2.GenerateHash())
	require.Equal(t, constants.DecDataID, testDecData.GetType())

	data, err := ReadDecData("")
	require.Equal(t, decData{Value: sdkTypes.ZeroDec()}, data)
	require.Nil(t, err)

	_, err = ReadDecData("testString")
	require.NotNil(t, err)

	data, err = ReadDecData("123")
	require.Equal(t, decData{Value: sdkTypes.NewDec(123)}, data)
	require.Nil(t, err)

	require.Equal(t, false, testDecData.Compare(NewStringData("")) == 0)
	require.Equal(t, true, testDecData.Compare(NewDecData(sdkTypes.NewDec(12))) == 0)
}
