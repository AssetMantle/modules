// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_MaintainerID_Methods(t *testing.T) {
	classificationID := baseIDs.NewStringID("classificationID")
	identityID := baseIDs.NewStringID("identityID")

	testMaintainerID := NewMaintainerID(classificationID, identityID).(maintainerID)
	require.NotPanics(t, func() {
		require.Equal(t, maintainerID{ClassificationID: classificationID, IdentityID: identityID}, testMaintainerID)
		require.Equal(t, strings.Join([]string{classificationID.String(), identityID.String()}, "."), testMaintainerID.String())
		require.Equal(t, false, testMaintainerID.IsPartial())
		require.Equal(t, true, maintainerID{ClassificationID: baseIDs.NewStringID(""), IdentityID: baseIDs.NewStringID("")}.IsPartial())
		require.Equal(t, true, testMaintainerID.Equals(testMaintainerID))
		require.Equal(t, false, testMaintainerID.Equals(maintainerID{ClassificationID: baseIDs.NewStringID(""), IdentityID: baseIDs.NewStringID("")}))
		require.Equal(t, true, testMaintainerID.Equals(testMaintainerID))
		require.Equal(t, false, testMaintainerID.Equals(maintainerID{ClassificationID: baseIDs.NewStringID(""), IdentityID: baseIDs.NewStringID("")}))
		require.Equal(t, false, testMaintainerID.Equals(nil))
		require.Equal(t, testMaintainerID, FromID(testMaintainerID))
		require.Equal(t, maintainerID{ClassificationID: baseIDs.NewStringID(""), IdentityID: baseIDs.NewStringID("")}, FromID(baseIDs.NewStringID("")))
		require.Equal(t, testMaintainerID, readMaintainerID(testMaintainerID.String()))
	})

}
