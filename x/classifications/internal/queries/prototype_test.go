// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/classifications/internal/queries/classification"
)

func TestPrototype(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, Prototype().Get("classification").GetName(), baseHelpers.NewQueries(
			classification.Query,
		).Get("classification").GetName())
	})
}
