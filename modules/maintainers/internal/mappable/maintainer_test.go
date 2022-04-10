// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	"github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	qualifiedMappables "github.com/AssetMantle/modules/schema/mappables/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Maintainer_Methods(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	identityID := baseIDs.NewID("identityID")
	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID"), base.NewStringData("ImmutableData")))
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID"), base.NewStringData("MutableData")))

	testMaintainerID := key.NewMaintainerID(classificationID, identityID)

	testMaintainer := NewMaintainer(testMaintainerID, nil, mutableProperties).(maintainer)

	require.Equal(t, maintainer{Document: qualifiedMappables.Document{ID: testMaintainerID, ClassificationID: classificationID, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}}, testMaintainer)
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
