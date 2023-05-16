// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/queries/maintainer"
	"github.com/stretchr/testify/require"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().GetQuery("maintainer").GetName(), baseHelpers.NewQueries(
			maintainer.Query,
		).GetQuery("maintainer").GetName())
	})
}
