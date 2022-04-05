// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

func Test_HeightData(t *testing.T) {
	heightValue := base.NewHeight(123)
	testHeightData := NewHeightData(heightValue)
	testHeightData2 := NewHeightData(base.NewHeight(0))

	require.Equal(t, "123", testHeightData.String())
	require.Equal(t, base.NewID(meta.Hash("123")), testHeightData.GenerateHashID())
	require.Equal(t, base.NewID(""), testHeightData2.GenerateHashID())
	require.Equal(t, HeightDataID, testHeightData.GetTypeID())

	dataAsString, err := testHeightData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsHeight, err := testHeightData.AsHeight()
	require.Equal(t, heightValue, dataAsHeight)
	require.Equal(t, nil, err)

	dataAsDec, err := testHeightData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, err)

	require.Equal(t, heightValue, testHeightData.Get())

	require.Equal(t, false, testHeightData.Compare(NewStringData("")) == 0)
	require.Equal(t, true, testHeightData.Compare(NewHeightData(base.NewHeight(123))) == 0)
}
