// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package splits

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/splits/auxiliaries"
	"github.com/AssetMantle/modules/x/splits/block"
	"github.com/AssetMantle/modules/x/splits/constants"
	"github.com/AssetMantle/modules/x/splits/genesis"
	"github.com/AssetMantle/modules/x/splits/invariants"
	"github.com/AssetMantle/modules/x/splits/mapper"
	"github.com/AssetMantle/modules/x/splits/migrations"
	"github.com/AssetMantle/modules/x/splits/parameters"
	"github.com/AssetMantle/modules/x/splits/queries"
	"github.com/AssetMantle/modules/x/splits/simulator"
	"github.com/AssetMantle/modules/x/splits/transactions"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		constants.ModuleName,
		constants.ModuleConsensusVersion,
		auxiliaries.Prototype,
		block.Prototype,
		genesis.Prototype,
		invariants.Prototype,
		mapper.Prototype,
		migrations.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
	)
}
