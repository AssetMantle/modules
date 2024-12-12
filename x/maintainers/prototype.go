// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainers

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries"
	"github.com/AssetMantle/modules/x/maintainers/block"
	"github.com/AssetMantle/modules/x/maintainers/constants"
	"github.com/AssetMantle/modules/x/maintainers/genesis"
	"github.com/AssetMantle/modules/x/maintainers/invariants"
	"github.com/AssetMantle/modules/x/maintainers/mapper"
	"github.com/AssetMantle/modules/x/maintainers/migrations"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
	"github.com/AssetMantle/modules/x/maintainers/queries"
	"github.com/AssetMantle/modules/x/maintainers/simulator"
	"github.com/AssetMantle/modules/x/maintainers/transactions"
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
