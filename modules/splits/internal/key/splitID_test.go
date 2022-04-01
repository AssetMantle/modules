// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_SplitID_Methods(t *testing.T) {
	ownerID := base.NewID("ownerID")
	ownableID := base.NewID("ownableID")

	testSplitID := NewSplitID(ownerID, ownableID).(splitID)
	testSplitID2 := NewSplitID(base.NewID(""), base.NewID("")).(splitID)
	require.NotPanics(t, func() {
		require.Equal(t, strings.Join([]string{ownerID.String(), ownableID.String()}, constants.SecondOrderCompositeIDSeparator), testSplitID.String())
		require.Equal(t, true, testSplitID.Equals(testSplitID))
		require.Equal(t, false, testSplitID.Equals(testSplitID2))
		require.Equal(t, false, testSplitID.IsPartial())
		require.Equal(t, true, testSplitID2.IsPartial())

		require.Equal(t, true, testSplitID.Equals(testSplitID))
		require.Equal(t, false, testSplitID.Equals(testSplitID2))
		require.Equal(t, false, testSplitID.Equals(nil))
		require.Equal(t, testSplitID, FromID(testSplitID))
		require.Equal(t, testSplitID2, FromID(base.NewID("")))
		require.Equal(t, splitID{OwnerID: base.NewID("ID1"), OwnableID: base.NewID("ID2")}, readSplitID("ID1*ID2"))
	})
}
