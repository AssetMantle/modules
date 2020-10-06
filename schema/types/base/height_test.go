/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHeight(t *testing.T) {
	height := height{Height: 10}
	heightByFunction := NewHeight(10)
	require.Equal(t, height.Height, heightByFunction.Get())
	require.Equal(t, true, heightByFunction.IsGreaterThan(NewHeight(1)))
	require.Equal(t, false, heightByFunction.IsGreaterThan(NewHeight(10)))
	require.Equal(t, false, heightByFunction.IsGreaterThan(NewHeight(20)))

}
