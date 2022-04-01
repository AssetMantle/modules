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

func Test_HeightData(t *testing.T) {
	heightValue := NewHeight(123)
	testHeightData := NewHeightData(heightValue)
	testHeightData2 := NewHeightData(NewHeight(0))

	require.Equal(t, "123", testHeightData.String())
	require.Equal(t, NewID(meta.Hash("123")), testHeightData.GenerateHashID())
	require.Equal(t, NewID(""), testHeightData2.GenerateHashID())
	require.Equal(t, heightDataID, testHeightData.GetTypeID())

	dataAsString, err := testHeightData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsHeight, err := testHeightData.AsHeight()
	require.Equal(t, heightValue, dataAsHeight)
	require.Equal(t, nil, err)

	dataAsDec, err := testHeightData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsID, err := testHeightData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, err)

	require.Equal(t, heightValue, testHeightData.Get())

	data, err := ReadHeightData("")
	require.Equal(t, heightData{Value: height{Value: 0}}, data)
	require.Nil(t, err)

	data, err = ReadHeightData("testString")
	require.Equal(t, nil, data)
	require.NotNil(t, err)

	data, err = ReadHeightData("123")
	require.Equal(t, heightData{Value: height{Value: 123}}, data)
	require.Nil(t, err)

	require.Equal(t, false, testHeightData.Compare(NewStringData("")) == 0)
	require.Equal(t, true, testHeightData.Compare(NewHeightData(NewHeight(123))) == 0)
}
