// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

func Test_StringData(t *testing.T) {
	value := "data"
	testStringData := NewStringData(value)
	testStringData2 := NewStringData("")

	require.Equal(t, value, testStringData.String())
	require.Equal(t, NewID(meta.Hash(value)), testStringData.GenerateHashID())
	require.Equal(t, NewID(""), testStringData2.GenerateHashID())
	require.Equal(t, stringDataID, testStringData.GetTypeID())
	require.Equal(t, testStringData.ZeroValue(), NewStringData(""))

	dataAsString, err := testStringData.AsString()
	require.Equal(t, value, dataAsString)
	require.Equal(t, nil, err)

	dataAsID, err := testStringData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsHeight, err := testStringData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsDec, err := testStringData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, err)

	require.Equal(t, value, testStringData.Get())

	data, err := ReadStringData("testString")
	require.Nil(t, err)
	require.Equal(t, stringData{Value: "testString"}.String(), data.String())

	require.Equal(t, false, testStringData.Compare(testStringData2) == 0)
	require.Equal(t, true, testStringData.Compare(testStringData) == 0)
	require.Panics(t, func() {
		require.Equal(t, false, testStringData.Compare(NewIDData(NewID("ID"))) == 0)
	})
}
