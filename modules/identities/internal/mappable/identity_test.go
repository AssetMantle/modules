// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_Identity_Methods(t *testing.T) {

	classificationID := baseIDs.NewStringID("classificationID")
	defaultImmutableProperties, _ := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	testIdentityID := key.NewIdentityID(classificationID, defaultImmutableProperties)
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData")))

	testIdentity := NewIdentity(testIdentityID, immutableProperties, mutableProperties)
	require.Equal(t, testIdentity, identity{Document: baseQualified.Document{ID: testIdentityID, Immutables: baseQualified.Immutables{PropertyList: immutableProperties}, Mutables: baseQualified.Mutables{Properties: mutableProperties}}})
	require.Equal(t, testIdentity.(identity).GetID(), testIdentityID)
	require.Equal(t, testIdentity.GetImmutablePropertyList(), immutableProperties)
	require.Equal(t, testIdentity.GetMutablePropertyList(), mutableProperties)
	require.Equal(t, testIdentity.GetKey(), testIdentityID)
}
