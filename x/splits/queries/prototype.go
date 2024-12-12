// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/queries/balances"
	"github.com/AssetMantle/modules/x/splits/queries/parameters"
	"github.com/AssetMantle/modules/x/splits/queries/split"
	"github.com/AssetMantle/modules/x/splits/queries/splits"
	"github.com/AssetMantle/modules/x/splits/queries/supply"
)

func Prototype() helpers.Queries {
	return baseHelpers.NewQueries(
		balances.Query,
		split.Query,
		splits.Query,
		supply.Query,
		parameters.Query,
	)
}
