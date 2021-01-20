/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeight(t *testing.T) {
	height := height{Value: 10}
	heightByFunction := NewHeight(10)
	require.Equal(t, height.Value, heightByFunction.Get())
	require.Equal(t, true, heightByFunction.IsGreaterThan(NewHeight(1)))
	require.Equal(t, false, heightByFunction.IsGreaterThan(NewHeight(10)))
	require.Equal(t, false, heightByFunction.IsGreaterThan(NewHeight(20)))

}
