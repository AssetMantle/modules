// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype(), baseHelpers.NewMapper(key.Prototype, mappable.Prototype))
	})
}
