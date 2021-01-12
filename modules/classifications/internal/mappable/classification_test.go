/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Classification_Methods(t *testing.T) {

	immutables := base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData")))))
	mutables := base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("MutableData")))))

	chainID := base.NewID("chainID")
	id := key.NewClassificationID(chainID, immutables, mutables)

	testClassification := NewClassification(id, immutables, mutables)
	require.Equal(t, classification{ID: id, ImmutableTraits: immutables, MutableTraits: mutables}, testClassification)
	require.Equal(t, immutables, testClassification.GetImmutables())
	require.Equal(t, mutables, testClassification.GetMutables())
	require.Equal(t, key.New(id), testClassification.GetKey())
	require.Equal(t, id, testClassification.(classification).GetID())
}
