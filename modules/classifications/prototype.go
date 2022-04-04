// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classifications

import (
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries"
	"github.com/AssetMantle/modules/modules/classifications/internal/block"
	"github.com/AssetMantle/modules/modules/classifications/internal/genesis"
	"github.com/AssetMantle/modules/modules/classifications/internal/mapper"
	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/modules/classifications/internal/parameters"
	"github.com/AssetMantle/modules/modules/classifications/internal/queries"
	"github.com/AssetMantle/modules/modules/classifications/internal/simulator"
	"github.com/AssetMantle/modules/modules/classifications/internal/transactions"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Module {
	return base.NewModule(
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
