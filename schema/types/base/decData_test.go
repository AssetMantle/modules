/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

func Test_DecData(t *testing.T) {

	decValue := sdkTypes.NewDec(12)
	testDecData := NewDecData(decValue)
	testDecData2 := NewDecData(sdkTypes.NewDec(0))

	require.Equal(t, decValue.String(), testDecData.String())
	require.Equal(t, NewID(meta.Hash(decValue.String())), testDecData.GenerateHashID())
	require.Equal(t, NewID(""), testDecData2.GenerateHashID())
	require.Equal(t, NewID("D"), testDecData.GetTypeID())

	dataAsString, err := testDecData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsDec, err := testDecData.AsDec()
	require.Equal(t, decValue, dataAsDec)
	require.Equal(t, nil, err)

	dataAsHeight, err := testDecData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsID, err := testDecData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, err)
	require.Equal(t, decValue, testDecData.Get())

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
