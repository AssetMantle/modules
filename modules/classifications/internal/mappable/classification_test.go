// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	qualifiedMappables "github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/qualified/base"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Classification_Methods(t *testing.T) {

	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewStringData("MutableData")))

	chainID := base.NewID("chainID")
	id := key.NewClassificationID(chainID, immutableProperties, mutableProperties)

	testClassification := NewClassification(id, immutableProperties, mutableProperties)
	require.Equal(t, classification{Document: qualifiedMappables.Document{ID: id, HasImmutables: base.HasImmutables{Properties: immutableProperties}, HasMutables: base.HasMutables{Properties: mutableProperties}}}, testClassification)
	require.Equal(t, immutableProperties, testClassification.GetImmutableProperties())
	require.Equal(t, mutableProperties, testClassification.GetMutableProperties())
	require.Equal(t, key.FromID(id), testClassification.GetKey())
	require.Equal(t, id, testClassification.(classification).GetID())
}
