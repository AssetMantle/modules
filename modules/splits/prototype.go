// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package splits

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/modules/splits/auxiliaries"
	"github.com/AssetMantle/modules/modules/splits/internal/block"
	"github.com/AssetMantle/modules/modules/splits/internal/genesis"
	"github.com/AssetMantle/modules/modules/splits/internal/invariants"
	"github.com/AssetMantle/modules/modules/splits/internal/mapper"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/modules/splits/internal/parameters"
	"github.com/AssetMantle/modules/modules/splits/internal/queries"
	"github.com/AssetMantle/modules/modules/splits/internal/simulator"
	"github.com/AssetMantle/modules/modules/splits/internal/transactions"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		module.Name,
		module.ConsensusVersion,
		auxiliaries.Prototype,
		block.Prototype,
		genesis.Prototype,
		invariants.Prototype,
		mapper.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
	)
}
