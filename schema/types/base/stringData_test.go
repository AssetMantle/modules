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

func Test_StringData(t *testing.T) {

	value := "data"
	testStringData := NewStringData(value)
	testStringData2 := NewStringData("")

	require.Equal(t, value, testStringData.String())
	require.Equal(t, NewID(meta.Hash(value)), testStringData.GenerateHashID())
	require.Equal(t, NewID(""), testStringData2.GenerateHashID())
	require.Equal(t, testStringData.GetTypeID(), NewID("S"))
	require.Equal(t, testStringData.ZeroValue(), NewStringData(""))

	dataAsString, Error := testStringData.AsString()
	require.Equal(t, value, dataAsString)
	require.Equal(t, nil, Error)

	dataAsID, Error := testStringData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsHeight, Error := testStringData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsDec, Error := testStringData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, Error)

	require.Equal(t, value, testStringData.Get())

	data, Error := ReadStringData("testString")
	require.Equal(t, stringData{Value: "testString"}.String(), data.String())
	require.Nil(t, Error)

	require.Equal(t, false, testStringData.Equal(testStringData2))
	require.Equal(t, true, testStringData.Equal(testStringData))
	require.Equal(t, false, testStringData.Equal(NewIDData(NewID("ID"))))
}
