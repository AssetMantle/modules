// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transaction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterCodec(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, RegisterCodec(nil), nil)
	})
}
