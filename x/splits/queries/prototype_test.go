// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/queries/ownable"
	"github.com/AssetMantle/modules/x/splits/queries/split"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().GetQuery("splits").GetName(), baseHelpers.NewQueries(
			split.Query,
			ownable.Query,
		).GetQuery("splits").GetName())
		require.Equal(t, Prototype().GetQuery("ownable").GetName(), baseHelpers.NewQueries(
			split.Query,
			ownable.Query,
		).GetQuery("ownable").GetName())
	})
}
