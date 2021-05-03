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

func Test_HeightData(t *testing.T) {

	heightValue := NewHeight(123)
	testHeightData := NewHeightData(heightValue)
	testHeightData2 := NewHeightData(NewHeight(0))

	require.Equal(t, "123", testHeightData.String())
	require.Equal(t, NewID(meta.Hash("123")), testHeightData.GenerateHashID())
	require.Equal(t, NewID(""), testHeightData2.GenerateHashID())
	require.Equal(t, NewID("H"), testHeightData.GetTypeID())

	dataAsString, Error := testHeightData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsHeight, Error := testHeightData.AsHeight()
	require.Equal(t, heightValue, dataAsHeight)
	require.Equal(t, nil, Error)

	dataAsDec, Error := testHeightData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsID, Error := testHeightData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, Error)

	require.Equal(t, heightValue, testHeightData.Get())

	data, Error := ReadHeightData("")
	require.Equal(t, heightData{Value: height{Value: 0}}, data)
	require.Nil(t, Error)

	data, Error = ReadHeightData("testString")
	require.Equal(t, nil, data)
	require.NotNil(t, Error)

	data, Error = ReadHeightData("123")
	require.Equal(t, heightData{Value: height{Value: 123}}, data)
	require.Nil(t, Error)

	require.Equal(t, false, testHeightData.Equal(NewStringData("")))
	require.Equal(t, true, testHeightData.Equal(NewHeightData(NewHeight(123))))

}
