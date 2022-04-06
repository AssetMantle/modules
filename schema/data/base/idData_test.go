// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

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

	require.Equal(t, true, NewIDData(base.NewID("identity2")).Compare(NewIDData(base.NewID("identity2"))) == 0)

	require.Panics(t, func() {
		require.Equal(t, false, testIDData.Compare(NewStringData("")) == 0)
	})
	require.Equal(t, true, testIDData.Compare(testIDData) == 0)

	require.Equal(t, "", testIDData.ZeroValue().String())
}
