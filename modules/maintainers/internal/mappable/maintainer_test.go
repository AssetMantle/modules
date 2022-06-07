// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_Maintainer_Methods(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	identityID := baseIDs.NewID("identityID")
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID"), base.NewStringData("ImmutableData")))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID"), base.NewStringData("MutableData")))

	testMaintainerID := key.NewMaintainerID(classificationID, identityID)

	testMaintainer := NewMaintainer(testMaintainerID, nil, mutableProperties).(maintainer)

	require.Equal(t, maintainer{Document: baseQualified.Document{ID: testMaintainerID, ClassificationID: classificationID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}}, testMaintainer)
	require.Equal(t, testMaintainerID, testMaintainer.GetID())
	require.Equal(t, classificationID, testMaintainer.GetClassificationID())
	require.Equal(t, identityID, testMaintainer.GetIdentityID())
	require.Equal(t, true, testMaintainer.CanAddMaintainer())
	require.Equal(t, true, testMaintainer.CanMutateMaintainer())
	require.Equal(t, true, testMaintainer.CanRemoveMaintainer())
	require.Equal(t, true, testMaintainer.MaintainsProperty(baseIDs.NewID("ID")))
	require.Equal(t, false, testMaintainer.MaintainsProperty(baseIDs.NewID("ID2")))
	require.Equal(t, testMaintainerID, testMaintainer.GetKey())
}
