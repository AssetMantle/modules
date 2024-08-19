// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/queries/balances"
	"github.com/AssetMantle/modules/x/splits/queries/split"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().GetQuery("splits").GetServicePath(), baseHelpers.NewQueries(
			split.Query,
			balances.Query,
		).GetQuery("splits").GetServicePath())
		require.Equal(t, Prototype().GetQuery("balances").GetServicePath(), baseHelpers.NewQueries(
			split.Query,
			balances.Query,
		).GetQuery("balances").GetServicePath())
	})
}
