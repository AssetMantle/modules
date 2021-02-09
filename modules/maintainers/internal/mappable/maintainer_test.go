/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Maintainer_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")
	properties := base.NewProperties(base.NewProperty(base.NewID("ID"), base.NewFact(base.NewStringData("MutableData"))))

	testMaintainerID := key.NewMaintainerID(classificationID, identityID)

	testMaintainer := NewMaintainer(testMaintainerID, properties, true, true, true).(maintainer)

	require.Equal(t, maintainer{ID: testMaintainerID, MaintainedProperties: properties, AddMaintainer: true, RemoveMaintainer: true, MutateMaintainer: true}, testMaintainer)
	require.Equal(t, testMaintainerID, testMaintainer.GetID())
	require.Equal(t, classificationID, testMaintainer.GetClassificationID())
	require.Equal(t, identityID, testMaintainer.GetIdentityID())
	require.Equal(t, true, testMaintainer.CanAddMaintainer())
	require.Equal(t, true, testMaintainer.CanMutateMaintainer())
	require.Equal(t, true, testMaintainer.CanRemoveMaintainer())
	require.Equal(t, true, testMaintainer.MaintainsProperty(base.NewID("ID")))
	require.Equal(t, false, testMaintainer.MaintainsProperty(base.NewID("ID2")))
	require.Equal(t, testMaintainerID, testMaintainer.GetKey())
}
