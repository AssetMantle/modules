// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package orders

import (
	"github.com/AssetMantle/modules/modules/orders/auxiliaries"
	"github.com/AssetMantle/modules/modules/orders/internal/block"
	"github.com/AssetMantle/modules/modules/orders/internal/genesis"
	"github.com/AssetMantle/modules/modules/orders/internal/mapper"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/orders/internal/parameters"
	"github.com/AssetMantle/modules/modules/orders/internal/queries"
	"github.com/AssetMantle/modules/modules/orders/internal/simulator"
	"github.com/AssetMantle/modules/modules/orders/internal/transactions"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		module.Name,
		auxiliaries.Prototype,
		genesis.Prototype,
		mapper.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
		block.Prototype,
	)
}
