// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/internal/queries/ownable"
	"github.com/AssetMantle/modules/x/splits/internal/queries/split"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().Get("splits").GetName(), baseHelpers.NewQueries(
			split.Query,
			ownable.Query,
		).Get("splits").GetName())
		require.Equal(t, Prototype().Get("ownable").GetName(), baseHelpers.NewQueries(
			split.Query,
			ownable.Query,
		).Get("ownable").GetName())
	})
}
