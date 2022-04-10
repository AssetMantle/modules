// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, block{}, Prototype())
}
