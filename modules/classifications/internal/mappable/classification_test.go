// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	qualifiedMappables "github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	baseQualified "github.com/persistenceOne/persistenceSDK/schema/qualified/base"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	baseTypes "github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Classification_Methods(t *testing.T) {

	immutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseTypes.NewID("ID1"), baseTypes.NewStringData("ImmutableData")))
	mutableProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseTypes.NewID("ID2"), baseTypes.NewStringData("MutableData")))

	chainID := baseTypes.NewID("chainID")
	id := key.NewClassificationID(chainID, immutableProperties, mutableProperties)

	testClassification := NewClassification(id, immutableProperties, mutableProperties)
	require.Equal(t, classification{Document: qualifiedMappables.Document{ID: id, HasImmutables: baseQualified.HasImmutables{Properties: immutableProperties}, HasMutables: baseQualified.HasMutables{Properties: mutableProperties}}}, testClassification)
	require.Equal(t, immutableProperties, testClassification.GetImmutableProperties())
	require.Equal(t, mutableProperties, testClassification.GetMutableProperties())
	require.Equal(t, key.FromID(id), testClassification.GetKey())
	require.Equal(t, id, testClassification.(classification).GetID())
}
