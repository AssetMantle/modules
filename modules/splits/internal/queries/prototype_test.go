// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/splits/internal/queries/ownable"
	"github.com/AssetMantle/modules/modules/splits/internal/queries/split"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().Get("splits").GetName(), base.NewQueries(
			split.Query,
			ownable.Query,
		).Get("splits").GetName())
		require.Equal(t, Prototype().Get("ownable").GetName(), base.NewQueries(
			split.Query,
			ownable.Query,
		).Get("ownable").GetName())
	})
}
