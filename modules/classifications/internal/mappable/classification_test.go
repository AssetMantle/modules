// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	qualifiedMappables "github.com/AssetMantle/modules/schema/mappables/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Classification_Methods(t *testing.T) {

	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), base.NewStringData("ImmutableData")))
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID2"), base.NewStringData("MutableData")))

	chainID := baseIDs.NewID("chainID")
	id := key.NewClassificationID(chainID, immutableProperties, mutableProperties)

	testClassification := NewClassification(id, immutableProperties, mutableProperties)
	require.Equal(t, classification{Document: qualifiedMappables.Document{ID: id, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}}, testClassification)
	require.Equal(t, immutableProperties, testClassification.GetImmutableProperties())
	require.Equal(t, mutableProperties, testClassification.GetMutableProperties())
	require.Equal(t, key.FromID(id), testClassification.GetKey())
	require.Equal(t, id, testClassification.(classification).GetID())
}
