// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	qualifiedMappables "github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	baseQualified "github.com/persistenceOne/persistenceSDK/schema/qualified/base"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	baseTypes "github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Maintainer_Methods(t *testing.T) {
	classificationID := baseTypes.NewID("classificationID")
	identityID := baseTypes.NewID("identityID")
	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseTypes.NewID("ID"), baseTypes.NewStringData("ImmutableData")))
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseTypes.NewID("ID"), baseTypes.NewStringData("MutableData")))

	testMaintainerID := key.NewMaintainerID(classificationID, identityID)

	testMaintainer := NewMaintainer(testMaintainerID, nil, mutableProperties).(maintainer)

	require.Equal(t, maintainer{Document: qualifiedMappables.Document{ID: testMaintainerID, ClassificationID: classificationID, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}}, testMaintainer)
	require.Equal(t, testMaintainerID, testMaintainer.GetID())
	require.Equal(t, classificationID, testMaintainer.GetClassificationID())
	require.Equal(t, identityID, testMaintainer.GetIdentityID())
	require.Equal(t, true, testMaintainer.CanAddMaintainer())
	require.Equal(t, true, testMaintainer.CanMutateMaintainer())
	require.Equal(t, true, testMaintainer.CanRemoveMaintainer())
	require.Equal(t, true, testMaintainer.MaintainsProperty(baseTypes.NewID("ID")))
	require.Equal(t, false, testMaintainer.MaintainsProperty(baseTypes.NewID("ID2")))
	require.Equal(t, testMaintainerID, testMaintainer.GetKey())
}
