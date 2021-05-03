/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
)

func Test_DecData(t *testing.T) {

	decValue := sdkTypes.NewDec(12)
	testDecData := NewDecData(decValue)
	testDecData2 := NewDecData(sdkTypes.NewDec(0))

	require.Equal(t, decValue.String(), testDecData.String())
	require.Equal(t, NewID(meta.Hash(decValue.String())), testDecData.GenerateHashID())
	require.Equal(t, NewID(""), testDecData2.GenerateHashID())
	require.Equal(t, NewID("D"), testDecData.GetTypeID())

	dataAsString, Error := testDecData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsDec, Error := testDecData.AsDec()
	require.Equal(t, decValue, dataAsDec)
	require.Equal(t, nil, Error)

	dataAsHeight, Error := testDecData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsID, Error := testDecData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, Error)
	require.Equal(t, decValue, testDecData.Get())

	data, Error := ReadDecData("")
	require.Equal(t, decData{Value: sdkTypes.ZeroDec()}, data)
	require.Nil(t, Error)

	data, Error = ReadDecData("testString")
	require.NotNil(t, Error)

	data, Error = ReadDecData("123")
	require.Equal(t, decData{Value: sdkTypes.NewDec(123)}, data)
	require.Nil(t, Error)

	require.Equal(t, false, testDecData.Equal(NewStringData("")))
	require.Equal(t, true, testDecData.Equal(NewDecData(sdkTypes.NewDec(12))))

}
