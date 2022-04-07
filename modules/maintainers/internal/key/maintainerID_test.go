// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_MaintainerID_Methods(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	identityID := baseIDs.NewID("identityID")

	testMaintainerID := NewMaintainerID(classificationID, identityID).(maintainerID)
	require.NotPanics(t, func() {
		require.Equal(t, maintainerID{ClassificationID: classificationID, IdentityID: identityID}, testMaintainerID)
		require.Equal(t, strings.Join([]string{classificationID.String(), identityID.String()}, constants.SecondOrderCompositeIDSeparator), testMaintainerID.String())
		require.Equal(t, false, testMaintainerID.IsPartial())
		require.Equal(t, true, maintainerID{ClassificationID: baseIDs.NewID(""), IdentityID: baseIDs.NewID("")}.IsPartial())
		require.Equal(t, true, testMaintainerID.Equals(testMaintainerID))
		require.Equal(t, false, testMaintainerID.Equals(maintainerID{ClassificationID: baseIDs.NewID(""), IdentityID: baseIDs.NewID("")}))
		require.Equal(t, true, testMaintainerID.Equals(testMaintainerID))
		require.Equal(t, false, testMaintainerID.Equals(maintainerID{ClassificationID: baseIDs.NewID(""), IdentityID: baseIDs.NewID("")}))
		require.Equal(t, false, testMaintainerID.Equals(nil))
		require.Equal(t, testMaintainerID, FromID(testMaintainerID))
		require.Equal(t, maintainerID{ClassificationID: baseIDs.NewID(""), IdentityID: baseIDs.NewID("")}, FromID(baseIDs.NewID("")))
		require.Equal(t, testMaintainerID, readMaintainerID(testMaintainerID.String()))
	})

}
