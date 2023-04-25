// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/AssetMantle/modules/x/orders/queries/order"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("orders").GetName(), baseHelpers.NewQueries(
		order.Query,
	).Get("orders").GetName())
}
