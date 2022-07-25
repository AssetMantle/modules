// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/data/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_HeightData(t *testing.T) {
	heightValue := baseTypes.NewHeight(123)
	testHeightData := NewHeightData(heightValue)

	require.Equal(t, "123", testHeightData.String())
	require.Equal(t, constants.HeightDataID, testHeightData.GetType())

	require.Equal(t, false, testHeightData.Compare(NewStringData("")) == 0)
	require.Equal(t, true, testHeightData.Compare(NewHeightData(baseTypes.NewHeight(123))) == 0)
}
