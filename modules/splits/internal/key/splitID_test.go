// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_SplitID_Methods(t *testing.T) {
	ownerID := baseIDs.NewID("ownerID")
	ownableID := baseIDs.NewID("ownableID")

	testSplitID := NewSplitID(ownerID, ownableID).(splitID)
	testSplitID2 := NewSplitID(baseIDs.NewID(""), baseIDs.NewID("")).(splitID)
	require.NotPanics(t, func() {
		require.Equal(t, strings.Join([]string{ownerID.String(), ownableID.String()}, "."), testSplitID.String())
		require.Equal(t, true, testSplitID.Equals(testSplitID))
		require.Equal(t, false, testSplitID.Equals(testSplitID2))
		require.Equal(t, false, testSplitID.IsPartial())
		require.Equal(t, true, testSplitID2.IsPartial())

		require.Equal(t, true, testSplitID.Equals(testSplitID))
		require.Equal(t, false, testSplitID.Equals(testSplitID2))
		require.Equal(t, false, testSplitID.Equals(nil))
		require.Equal(t, testSplitID, FromID(testSplitID))
		require.Equal(t, testSplitID2, FromID(baseIDs.NewID("")))
		require.Equal(t, splitID{OwnerID: baseIDs.NewID("ID1"), OwnableID: baseIDs.NewID("ID2")}, readSplitID("ID1*ID2"))
	})
}
