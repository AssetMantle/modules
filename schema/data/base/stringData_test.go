// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/data/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/string"
)

func Test_StringData(t *testing.T) {
	value := "data"
	testStringData := NewStringData(value)
	testStringData2 := NewStringData("")

	require.Equal(t, value, testStringData.String())
	require.Equal(t, baseIDs.NewStringID(string.Hash(value)), testStringData.GenerateHash())
	require.Equal(t, baseIDs.NewStringID(""), testStringData2.GenerateHash())
	require.Equal(t, constants.StringDataID, testStringData.GetType())
	require.Equal(t, testStringData.ZeroValue(), NewStringData(""))
	require.Equal(t, false, testStringData.Compare(testStringData2) == 0)
	require.Equal(t, true, testStringData.Compare(testStringData) == 0)
	require.Panics(t, func() {
		require.Equal(t, false, testStringData.Compare(NewIDData(baseIDs.NewStringID("ID"))) == 0)
	})
}
