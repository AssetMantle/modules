// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/AssetMantle/modules/x/splits/queries/ownable"
	"github.com/AssetMantle/modules/x/splits/queries/split"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
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
