// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/queries/order"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().GetQuery("orders").GetName(), baseHelpers.NewQueries(
		order.Query,
	).GetQuery("orders").GetName())
}
