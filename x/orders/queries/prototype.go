// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/queries/order"
	"github.com/AssetMantle/modules/x/orders/queries/orders"
	"github.com/AssetMantle/modules/x/orders/queries/parameters"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		order.Query,
		orders.Query,
		parameters.Query,
	)
}
