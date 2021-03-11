/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"

	"github.com/stretchr/testify/require"
)

func Test_Mutables(t *testing.T) {

	testProperty := base.NewProperty(base.NewID("ID"), base.NewFact(base.NewStringData("Data")))
	testProperties := base.NewProperties(testProperty)
	testMutables := HasMutables{testProperties}
	require.Equal(t, HasMutables{Properties: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.GetMutableProperties())
	mutatedTestProperty := base.NewProperty(base.NewID("ID"), base.NewFact(base.NewStringData("Data2")))
	require.Equal(t, HasMutables{Properties: base.NewProperties(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}
