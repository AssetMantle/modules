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
	testMutables := Mutables{testProperties}
	require.Equal(t, Mutables{Properties: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.GetMutables())
	mutatedTestProperty := base.NewProperty(base.NewID("ID"), base.NewFact(base.NewStringData("Data2")))
	require.Equal(t, Mutables{Properties: base.NewProperties(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}
