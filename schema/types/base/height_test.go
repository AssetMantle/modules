// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeight(t *testing.T) {
	h := height{Value: 10}
	heightByFunction := NewHeight(10)

	require.Equal(t, h.Value, heightByFunction.Get())
	require.Equal(t, true, heightByFunction.Compare(NewHeight(1)) > 0)
	require.Equal(t, false, heightByFunction.Compare(NewHeight(10)) > 0)
	require.Equal(t, false, heightByFunction.Compare(NewHeight(20)) > 0)
}
