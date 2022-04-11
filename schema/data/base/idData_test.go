// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

func Test_IDData(t *testing.T) {
	idValue := baseIDs.NewID("ID")
	testIDData := NewIDData(idValue)
	testIDData2 := NewIDData(baseIDs.NewID(""))

	require.Equal(t, "ID", testIDData.String())
	require.Equal(t, baseIDs.NewID(meta.Hash("ID")), testIDData.GenerateHash())
	require.Equal(t, baseIDs.NewID(""), testIDData2.GenerateHash())
	require.Equal(t, ids.IDDataID, testIDData.GetType())

	require.Equal(t, true, NewIDData(baseIDs.NewID("identity2")).Compare(NewIDData(baseIDs.NewID("identity2"))) == 0)

	require.Panics(t, func() {
		require.Equal(t, false, testIDData.Compare(NewStringData("")) == 0)
	})
	require.Equal(t, true, testIDData.Compare(testIDData) == 0)

	require.Equal(t, "", testIDData.ZeroValue().String())
}
