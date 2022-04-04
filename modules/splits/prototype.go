// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package splits

import (
	"github.com/AssetMantle/modules/modules/splits/auxiliaries"
	"github.com/AssetMantle/modules/modules/splits/internal/block"
	"github.com/AssetMantle/modules/modules/splits/internal/genesis"
	"github.com/AssetMantle/modules/modules/splits/internal/mapper"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/modules/splits/internal/queries"
	"github.com/AssetMantle/modules/modules/splits/internal/simulator"
	"github.com/AssetMantle/modules/modules/splits/internal/transactions"
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
