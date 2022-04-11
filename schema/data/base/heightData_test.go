// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

func Test_HeightData(t *testing.T) {
	heightValue := baseTypes.NewHeight(123)
	testHeightData := NewHeightData(heightValue)
	testHeightData2 := NewHeightData(baseTypes.NewHeight(0))

	require.Equal(t, "123", testHeightData.String())
	require.Equal(t, baseIDs.NewID(meta.Hash("123")), testHeightData.GenerateHash())
	require.Equal(t, baseIDs.NewID(""), testHeightData2.GenerateHash())
	require.Equal(t, ids.HeightDataID, testHeightData.GetType())

	require.Equal(t, false, testHeightData.Compare(NewStringData("")) == 0)
	require.Equal(t, true, testHeightData.Compare(NewHeightData(baseTypes.NewHeight(123))) == 0)
}
