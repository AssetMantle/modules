// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func Test_Classification_Methods(t *testing.T) {

	immutableProperties := baseLists.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData")))

	id := key.NewClassificationID(immutableProperties, mutableProperties)

	testClassification := NewClassification(immutableProperties, mutableProperties)
	require.Equal(t, classification{document: baseQualified.document{ID: id, Immutables: baseQualified.immutables{PropertyList: immutableProperties}, Mutables: baseQualified.mutables{PropertyList: mutableProperties}}}, testClassification)
	require.Equal(t, immutableProperties, testClassification.GetImmutables())
	require.Equal(t, mutableProperties, testClassification.GetMutables())
	require.Equal(t, key.FromID(id), testClassification.GetKey())
	require.Equal(t, id, testClassification.(classification).GetID())
}
