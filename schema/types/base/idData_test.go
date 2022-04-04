// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/utilities/meta"
)

func Test_IDData(t *testing.T) {
	idValue := NewID("ID")
	testIDData := NewIDData(idValue)
	testIDData2 := NewIDData(NewID(""))

	require.Equal(t, "ID", testIDData.String())
	require.Equal(t, NewID(meta.Hash("ID")), testIDData.GenerateHashID())
	require.Equal(t, NewID(""), testIDData2.GenerateHashID())
	require.Equal(t, idDataID, testIDData.GetTypeID())

	dataAsString, err := testIDData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsID, err := testIDData.AsID()
	require.Equal(t, idValue, dataAsID)
	require.Equal(t, nil, err)

	dataAsHeight, err := testIDData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsDec, err := testIDData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, err)

	require.Equal(t, idValue, testIDData.Get())

	data, err := ReadIDData("testString")
	require.Equal(t, idData{Value: id{IDString: "testString"}}, data)
	require.Nil(t, err)
	require.Equal(t, true, NewIDData(NewID("identity2")).Compare(NewIDData(NewID("identity2"))) == 0)

	require.Panics(t, func() {
		require.Equal(t, false, testIDData.Compare(NewStringData("")) == 0)
	})
	require.Equal(t, true, testIDData.Compare(testIDData) == 0)

	require.Equal(t, "", testIDData.ZeroValue().String())
}
