// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/x/orders/internal/queries/order"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("orders").GetName(), baseHelpers.NewQueries(
		order.Query,
	).Get("orders").GetName())
}
