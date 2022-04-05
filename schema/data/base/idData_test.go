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

func Test_IDData(t *testing.T) {
	idValue := base.NewID("ID")
	testIDData := NewIDData(idValue)
	testIDData2 := NewIDData(base.NewID(""))

	require.Equal(t, "ID", testIDData.String())
	require.Equal(t, base.NewID(meta.Hash("ID")), testIDData.GenerateHashID())
	require.Equal(t, base.NewID(""), testIDData2.GenerateHashID())
	require.Equal(t, IDDataID, testIDData.GetTypeID())

	dataAsString, err := testIDData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsID, err := testIDData.AsID()
	require.Equal(t, idValue, dataAsID)
	require.Equal(t, nil, err)

	dataAsDec, err := testIDData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, err)

	require.Equal(t, idValue, testIDData.Get())

	require.Equal(t, true, NewIDData(base.NewID("identity2")).Compare(NewIDData(base.NewID("identity2"))) == 0)

	require.Panics(t, func() {
		require.Equal(t, false, testIDData.Compare(NewStringData("")) == 0)
	})
	require.Equal(t, true, testIDData.Compare(testIDData) == 0)

	require.Equal(t, "", testIDData.ZeroValue().String())
}
