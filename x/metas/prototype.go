// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/metas/auxiliaries"
	"github.com/AssetMantle/modules/x/metas/block"
	"github.com/AssetMantle/modules/x/metas/constants"
	"github.com/AssetMantle/modules/x/metas/genesis"
	"github.com/AssetMantle/modules/x/metas/invariants"
	"github.com/AssetMantle/modules/x/metas/mapper"
	"github.com/AssetMantle/modules/x/metas/migrations"
	"github.com/AssetMantle/modules/x/metas/parameters"
	"github.com/AssetMantle/modules/x/metas/queries"
	"github.com/AssetMantle/modules/x/metas/simulator"
	"github.com/AssetMantle/modules/x/metas/transactions"
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
