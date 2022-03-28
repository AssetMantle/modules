/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_MaintainerID_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")

	testMaintainerID := NewMaintainerID(classificationID, identityID).(MaintainerID)
	require.NotPanics(t, func() {
		require.Equal(t, MaintainerID{ClassificationID: classificationID, IdentityID: identityID}, testMaintainerID)
		require.Equal(t, strings.Join([]string{classificationID.String(), identityID.String()}, constants.SecondOrderCompositeIDSeparator), testMaintainerID.String())
		require.Equal(t, false, testMaintainerID.IsPartial())
		require.Equal(t, true, MaintainerID{ClassificationID: base.NewID(""), IdentityID: base.NewID("")}.IsPartial())
		require.Equal(t, true, testMaintainerID.Equals(testMaintainerID))
		require.Equal(t, false, testMaintainerID.Equals(MaintainerID{ClassificationID: base.NewID(""), IdentityID: base.NewID("")}))
		require.Equal(t, true, testMaintainerID.Equals(testMaintainerID))
		require.Equal(t, false, testMaintainerID.Equals(MaintainerID{ClassificationID: base.NewID(""), IdentityID: base.NewID("")}))
		require.Equal(t, false, testMaintainerID.Equals(nil))
		require.Equal(t, testMaintainerID, FromID(&testMaintainerID))
		require.Equal(t, MaintainerID{ClassificationID: base.NewID(""), IdentityID: base.NewID("")}, FromID(base.NewID("")))
		require.Equal(t, testMaintainerID, readMaintainerID(testMaintainerID.String()))
	})

}
